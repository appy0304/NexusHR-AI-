package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// --- YOUR API KEYS ---
var GroqAPIKey = os.Getenv("GROQ_API_KEY")
var HuggingFaceToken = os.Getenv("HUGGINGFACE_TOKEN")
var PineconeAPIKey = os.Getenv("PINECONE_API_KEY")
const PineconeHostURL = "https://emp-management-zp59a8y.svc.aped-4627-b74a.pinecone.io"

// --- AI MODEL URLS ---
const HFEmbeddingURL = "https://api-inference.huggingface.co/pipeline/feature-extraction/sentence-transformers/all-MiniLM-L6-v2"
const GroqURL = "https://api.groq.com/openai/v1/chat/completions"

// ==========================================
// 1. GET EMBEDDINGS (Hugging Face)
// ==========================================
func getEmbedding(text string) ([]float32, error) {
	reqBody, _ := json.Marshal(map[string]string{"inputs": text})
	req, _ := http.NewRequest("POST", HFEmbeddingURL, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+HuggingFaceToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var embedding []float32
	if err := json.Unmarshal(bodyBytes, &embedding); err != nil {
		return nil, fmt.Errorf("failed to parse embedding: %v. Response: %s", err, string(bodyBytes))
	}
	return embedding, nil
}

// ==========================================
// 2. UPLOAD TO PINECONE
// ==========================================
func uploadToPinecone(documentID string, text string) error {
	vector, err := getEmbedding(text)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"vectors": []map[string]interface{}{
			{
				"id":     documentID,
				"values": vector,
				"metadata": map[string]string{
					"text": text,
				},
			},
		},
	}

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", PineconeHostURL+"/vectors/upsert", bytes.NewBuffer(reqBody))
	req.Header.Set("Api-Key", PineconeAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// ==========================================
// 3. SEARCH PINECONE
// ==========================================
func searchPinecone(question string) (string, error) {
	questionVector, err := getEmbedding(question)
	if err != nil {
		return "", err
	}

	payload := map[string]interface{}{
		"vector":          questionVector,
		"topK":            2,
		"includeMetadata": true,
	}

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", PineconeHostURL+"/query", bytes.NewBuffer(reqBody))
	req.Header.Set("Api-Key", PineconeAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	matches, ok := result["matches"].([]interface{})
	if !ok || len(matches) == 0 {
		return "No HR documents found.", nil
	}

	combinedContext := ""
	for _, match := range matches {
		matchMap := match.(map[string]interface{})
		metadata := matchMap["metadata"].(map[string]interface{})
		text := metadata["text"].(string)
		combinedContext += text + "\n---\n"
	}

	return combinedContext, nil
}

// ==========================================
// 4. ASK GROQ LLM
// ==========================================
func askGroq(question string, context string) (string, error) {
	systemPrompt := fmt.Sprintf(`You are a helpful HR Assistant. Answer the user's question using ONLY the provided HR context. If the answer is not in the context, politely say "I don't know based on the HR policy."
Context:
%s`, context)

	payload := map[string]interface{}{
		"model": "llama3-8b-8192",
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": question},
		},
	}

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", GroqURL, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+GroqAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	choices := result["choices"].([]interface{})
	firstChoice := choices[0].(map[string]interface{})
	message := firstChoice["message"].(map[string]interface{})
	return message["content"].(string), nil
}

// ==========================================
// 5. SERVER ENDPOINT HANDLERS
// ==========================================

func handleUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	var reqData struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}
	json.NewDecoder(r.Body).Decode(&reqData)

	err := uploadToPinecone(reqData.ID, reqData.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, `{"status": "success", "message": "Document uploaded to Pinecone!"}`)
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	var reqData struct {
		Question string `json:"question"`
	}
	json.NewDecoder(r.Body).Decode(&reqData)

	context, err := searchPinecone(reqData.Question)
	if err != nil {
		http.Error(w, "Failed to search vectors: "+err.Error(), http.StatusInternalServerError)
		return
	}

	answer, err := askGroq(reqData.Question, context)
	if err != nil {
		http.Error(w, "Failed to ask AI: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"answer": answer}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

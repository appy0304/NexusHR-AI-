package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"simple-go-api/dao"
)

// --- OpenAI Types ---

type embeddingReq struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}
type embeddingRes struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}
type chatMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type chatReq struct {
	Model    string    `json:"model"`
	Messages []chatMsg `json:"messages"`
}
type chatRes struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// --- Pinecone Types ---

type pineconeQueryReq struct {
	Vector          []float32 `json:"vector"`
	TopK            int       `json:"topK"`
	IncludeMetadata bool      `json:"includeMetadata"`
}
type pineconeQueryRes struct {
	Matches []struct {
		Score    float32 `json:"score"`
		Metadata struct {
			Text   string `json:"text"`
			Source string `json:"source"`
		} `json:"metadata"`
	} `json:"matches"`
}

// openaiCall is a small helper to POST to OpenAI and decode response
func openaiCall(path string, body interface{}, out interface{}) error {
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/"+path, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	resp, err := (&http.Client{Timeout: 30 * time.Second}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("openai %s error %d: %s", path, resp.StatusCode, string(b))
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// AskRAG is the core RAG function: embed → search → answer → save
func AskRAG(query, userID string) (string, []string, error) {
	// Step 1: Generate embedding for the user query
	var embRes embeddingRes
	err := openaiCall("embeddings", embeddingReq{
		Input: []string{query},
		Model: "text-embedding-3-small",
	}, &embRes)
	if err != nil || len(embRes.Data) == 0 {
		return "", nil, fmt.Errorf("embedding failed: %v", err)
	}
	vector := embRes.Data[0].Embedding

	// Step 2: Search Pinecone for relevant document chunks
	retrievedContext, sources := searchPinecone(vector)

	// Step 3: Build strict RAG prompt + call GPT-4o
	systemPrompt := `You are an HR assistant. Answer the question using ONLY the provided context.
If the answer is not in the context, say "I don't have information about that in our documents."
Be concise.

Context:
` + retrievedContext

	var chatResp chatRes
	err = openaiCall("chat/completions", chatReq{
		Model: "gpt-4o",
		Messages: []chatMsg{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: query},
		},
	}, &chatResp)
	if err != nil || len(chatResp.Choices) == 0 {
		return "", nil, fmt.Errorf("chat completion failed: %v", err)
	}
	answer := chatResp.Choices[0].Message.Content

	// Step 4: Save to MongoDB
	_ = dao.SaveChat(&dao.ChatHistory{
		UserID:  userID,
		Query:   query,
		Answer:  answer,
		Sources: sources,
	})

	return answer, sources, nil
}

// searchPinecone queries the Pinecone index for top-3 matching document chunks
func searchPinecone(vector []float32) (string, []string) {
	pineconeHost := os.Getenv("PINECONE_HOST") // e.g., "https://hr-index-abc123.svc.us-east1-gcp.pinecone.io"
	apiKey := os.Getenv("PINECONE_API_KEY")

	body, _ := json.Marshal(pineconeQueryReq{
		Vector:          vector,
		TopK:            3,
		IncludeMetadata: true,
	})

	req, _ := http.NewRequest("POST", pineconeHost+"/query", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", apiKey)

	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "No context found.", nil
	}
	defer resp.Body.Close()

	var result pineconeQueryRes
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "No context found.", nil
	}

	var contextText string
	var sources []string
	seen := map[string]bool{}
	for _, m := range result.Matches {
		contextText += m.Metadata.Text + "\n\n"
		if !seen[m.Metadata.Source] {
			sources = append(sources, m.Metadata.Source)
			seen[m.Metadata.Source] = true
		}
	}

	if contextText == "" {
		return "No context found.", nil
	}
	return contextText, sources
}

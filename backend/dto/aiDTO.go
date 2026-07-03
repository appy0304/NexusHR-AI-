package dto

// AskRequest is the JSON body for POST /api/v1/ai/ask
type AskRequest struct {
	Query  string `json:"query" binding:"required"`
	UserID string `json:"userId"`
}

// AskResponse is the JSON response from the AI endpoint
type AskResponse struct {
	Answer  string   `json:"answer"`
	Sources []string `json:"sources"`
}

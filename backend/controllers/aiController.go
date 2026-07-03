package controllers

import (
	"net/http"

	"simple-go-api/dto"
	"simple-go-api/services"

	"github.com/gin-gonic/gin"
)

// AskAI handles POST /api/v1/ai/ask
func AskAI(c *gin.Context) {
	var req dto.AskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid request", err, c.GetString("requestId")))
		return
	}

	answer, sources, err := services.AskRAG(req.Query, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail("AI processing failed", err, c.GetString("requestId")))
		return
	}

	c.JSON(http.StatusOK, dto.Success("AI response", dto.AskResponse{
		Answer:  answer,
		Sources: sources,
	}, c.GetString("requestId")))
}

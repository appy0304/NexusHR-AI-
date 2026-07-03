package routes

import (
	"simple-go-api/controllers"
	"simple-go-api/middleware"

	"github.com/gin-gonic/gin"
)

// AIRoutes sets up AI-related routes
func AIRoutes(router *gin.Engine) {
	ai := router.Group("/api/v1/ai")
	ai.Use(middleware.JWTAuth())
	{
		ai.POST("/ask", controllers.AskAI)
	}
}

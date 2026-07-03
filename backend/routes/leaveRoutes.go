package routes

import (
	"simple-go-api/controllers"

	"github.com/gin-gonic/gin"
)

// LeaveRoutes sets up all leave-related routes
func LeaveRoutes(router *gin.Engine) {
	leaves := router.Group("/api/v1/leaves")
	// All leave routes require authentication (JWT middleware applied globally or here)
	// leaves.Use(middleware.JWTAuth())
	{
		leaves.POST("", controllers.CreateLeave)
		leaves.GET("", controllers.GetLeaves)
		leaves.GET("/:id", controllers.GetLeave)
		leaves.PUT("/:id", controllers.UpdateLeave)
		leaves.GET("/balance", controllers.GetLeaveBalance)
	}
}
package routes

import (
	"simple-go-api/controllers"
	"simple-go-api/middleware"

	"github.com/gin-gonic/gin"
)

// EmployeeRoutes sets up all employee-related routes
func EmployeeRoutes(router *gin.Engine) {
	// All employee routes require authentication
	employees := router.Group("/api/v1/employees")
	employees.Use(middleware.JWTAuth())
	{
		employees.POST("", controllers.CreateEmployee)
		employees.GET("", controllers.GetEmployees)
		employees.GET("/:id", controllers.GetEmployee)
		employees.PUT("/:id", controllers.UpdateEmployee)
		employees.DELETE("/:id", controllers.DeleteEmployee)
	}
}

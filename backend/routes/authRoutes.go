package routes

import (
	"simple-go-api/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up all authentication-related routes
func AuthRoutes(router *gin.Engine) {
	// Public auth routes (no authentication required)
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/refresh", controllers.Refresh)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/create-admin", controllers.CreateAdmin) // One-time setup
	}

	// Protected routes example (uncomment when you add protected endpoints)
	// protected := router.Group("/api/v1")
	// protected.Use(middleware.JWTAuth())
	// {
	//     // HR Admin and Super Admin only
	//     hr := protected.Group("/hr")
	//     hr.Use(middleware.RequireRole("super_admin", "hr_admin"))
	//     {
	//         // HR-specific endpoints
	//     }

	//     // Manager and above
	//     manager := protected.Group("/manager")
	//     manager.Use(middleware.RequireRole("super_admin", "hr_admin", "manager"))
	//     {
	//         // Manager-specific endpoints
	//     }
	// }
}

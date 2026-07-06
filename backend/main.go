package main

// @title Employee Management API
// @version 2.0
// @description This is an enterprise-grade AI-Powered Employee Management Platform
// @host localhost:8080
// @BasePath /
// @schemes http https

import (
	"os"
	_ "simple-go-api/docs"

	"simple-go-api/config"
	"simple-go-api/middleware"
	"simple-go-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	// Apply middleware
	router.Use(middleware.RequestID())
	router.Use(middleware.StructuredLogger())
	router.Use(middleware.CORS())
	router.Use(middleware.PaginationParser())
	// router.Use(middleware.RateLimiter(100)) // Uncomment for production

	// Routes
	routes.AuthRoutes(router)
	routes.EmployeeRoutes(router)
	routes.LeaveRoutes(router)
	routes.AIRoutes(router)
	router.POST("/api/v1/upload", gin.WrapF(handleUpload))
    router.POST("/api/v1/chat", gin.WrapF(handleChat))

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "Employee Management Platform",
		})
	})

	router.Use(middleware.PrometheusMetrics())
	router.GET("/metrics", middleware.MetricsHandler())

	// router.Run(":8080")
	 port := os.Getenv("PORT")
 if port == "" {
     port = "8080"
 }
 router.Run(":" + port)
}

// package main

// // @title Simple Go API
// // @version 1.0
// // @description This is a simple Go CRUD API
// // @host localhost:8080
// // @BasePath /

// import (
// 	_ "simple-go-api/docs"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"

// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"

// 	"simple-go-api/config"
// 	"simple-go-api/routes"
// )

// func main() {

// 	config.ConnectDB()

// 	router := gin.Default()
// 	   router.Use(cors.New(cors.Config{
//         AllowOrigins:     []string{"http://localhost:5173"},
//         AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//         AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
//         AllowCredentials: true,
//     }))

// 	routes.UserRoutes(router)

// 	router.GET("/swagger/*any",
// 		ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	router.Run(":8080")
// }

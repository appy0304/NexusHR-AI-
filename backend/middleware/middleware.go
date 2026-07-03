package middleware

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// generateRequestID creates a unique request ID for tracing
func generateRequestID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// RequestID middleware adds a unique ID to every request
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := generateRequestID()
		c.Set("requestId", requestID)
		c.Header("X-Request-ID", requestID)
		c.Header("X-Correlation-ID", requestID)
		c.Next()
	}
}

// StructuredLogger logs each request with consistent format
func StructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Log after response
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Structured log output
		fmt.Printf("[REQUEST] id=%s method=%s path=%s status=%d latency=%v client=%s\n",
			c.GetString("requestId"),
			c.Request.Method,
			path,
			statusCode,
			latency,
			c.ClientIP(),
		)
	}
}

// RateLimiter simple rate limiting middleware (can be replaced with Redis-based)
func RateLimiter(maxRequests int) gin.HandlerFunc {
	requests := make(map[string]int)
	// window := time.Minute

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		// Clean old entries
		for ip := range requests {
			if time.Since(getLastRequestTime(ip)).Minutes() > 1 {
				delete(requests, ip)
			}
		}

		if requests[clientIP] >= maxRequests {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Rate limit exceeded. Try again later.",
				"error":   "too many requests",
			})
			c.Abort()
			return
		}

		requests[clientIP]++
		c.Next()
	}
}

func getLastRequestTime(ip string) time.Time {
	return time.Now()
}

// CORS middleware (replaces the inline CORS in main.go)
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Request-ID")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// PaginationParser parses and validates pagination query parameters
func PaginationParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("pageSize", "20")
		sort := c.DefaultQuery("sort", "createdAt")
		order := c.DefaultQuery("order", "desc")

		p, err := strconv.Atoi(page)
		if err != nil || p < 1 {
			p = 1
		}

		ps, err := strconv.Atoi(pageSize)
		if err != nil || ps < 1 || ps > 100 {
			ps = 20
		}

		c.Set("page", p)
		c.Set("pageSize", ps)
		c.Set("sort", sort)
		c.Set("order", order)

		c.Next()
	}
}

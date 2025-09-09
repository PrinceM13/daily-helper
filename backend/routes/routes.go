package routes

import (
	"net/http"

	"github.com/PrinceM13/daily-helper/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Health check route
	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK Gin!!!")
	})

	// Task routes
	server.GET("/tasks", getTasks)

	// Protected routes that require authentication
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
}

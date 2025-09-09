package main

import (
	"fmt"
	"os"

	"github.com/PrinceM13/daily-helper/internal/handler"
	"github.com/PrinceM13/daily-helper/internal/repository"
	"github.com/PrinceM13/daily-helper/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Health check
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	healthHandler := handler.NewHealthHandler("1.0.0", env)
	healthHandler.RegisterRoutes(server)

	// Task management
	taskRepo := repository.NewTaskRepo()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(server)

	fmt.Println("Server listening on :" + port)
	server.Run(":" + port)
}

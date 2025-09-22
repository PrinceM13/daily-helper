package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PrinceM13/daily-helper/internal/adapters/handler"
	"github.com/PrinceM13/daily-helper/internal/adapters/repository"
	"github.com/PrinceM13/daily-helper/internal/application"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Health check
	healthHandler := handler.NewHealthHandler("1.0.0", env)
	healthHandler.RegisterRoutes(server)

	// Task management
	taskRepo, err := repository.NewTaskRepo()
	if err != nil {
		log.Fatalf("Failed to create task repository: %v", err)
	}
	taskService := application.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(server)

	fmt.Println("Server listening on :" + port)
	server.Run(":" + port)
}

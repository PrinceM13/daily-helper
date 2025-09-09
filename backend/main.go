package main

import (
	"fmt"
	"os"

	"github.com/PrinceM13/daily-helper/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Println("Server listening on :" + port)
	server.Run(":" + port)
}

package routes

import (
	"net/http"

	"github.com/PrinceM13/daily-helper/models"
	"github.com/gin-gonic/gin"
)

func getTasks(c *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	limit := 10
	page := 1
	totalCount := len(tasks)
	totalPages := (totalCount + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"data":       tasks,
		"page":       page,
		"limit":      limit,
		"totalCount": totalCount,
		"totalPages": totalPages,
	})
}

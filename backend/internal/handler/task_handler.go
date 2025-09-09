package handler

import (
	"net/http"
	"strconv"

	"github.com/PrinceM13/daily-helper/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{Service: s}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	{
		tasks.GET("", h.GetTasks)
		tasks.GET("/:id", h.GetTaskByID)
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	tasks, totalCount, err := h.Service.GetTasks(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := (totalCount + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"data":       tasks,
		"page":       page,
		"limit":      limit,
		"totalCount": totalCount,
		"totalPages": totalPages,
	})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	// bind JSON and call service.UpdateTask...
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

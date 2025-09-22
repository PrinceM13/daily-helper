package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	startTime time.Time
	version   string
	env       string
}

func NewHealthHandler(version, env string) *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
		version:   version,
		env:       env,
	}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	uptime := time.Since(h.startTime).String()
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": h.version,
		"uptime":  uptime,
		"env":     h.env,
		"time":    time.Now().UTC(),
	})
}

func (h *HealthHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/health", h.HealthCheck)
}

package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthcheckHandlers struct {
	logger zap.Logger
}

func NewHealthCheckHandlers(logger zap.Logger) *HealthcheckHandlers {
	return &HealthcheckHandlers{
		logger: logger,
	}
}

func (h *HealthcheckHandlers) HealthCheck(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *HealthcheckHandlers) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Readness OK",
	})
}

func (h *HealthcheckHandlers) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Liveness OK",
	})
}

func (h *HealthcheckHandlers) TimeoutTest(c *gin.Context) {

	time.Sleep(20 * time.Second)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

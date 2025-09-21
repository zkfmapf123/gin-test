package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthcheckHandlers struct {
	logger zap.Logger
	job    chan<- func()
}

func NewHealthCheckHandlers(logger zap.Logger, job chan<- func()) *HealthcheckHandlers {
	return &HealthcheckHandlers{
		logger: logger,
		job:    job,
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

func (h *HealthcheckHandlers) TestWorker(c *gin.Context) {

	h.job <- func() {
		fmt.Println("hello world")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "test worker",
	})
}

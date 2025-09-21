package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/dispatcher/internal/validate"
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

// Timeout test
func (h *HealthcheckHandlers) TimeoutTest(c *gin.Context) {

	time.Sleep(20 * time.Second)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// Disptacher Test
func (h *HealthcheckHandlers) TestWorker(c *gin.Context) {

	h.job <- func() {
		fmt.Println("hello world")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "test worker",
	})
}

type TestObj struct {
	Name string `json:"name" binding:"required,min=2,max=50"`
	Age  int    `json:"age" binding:"required"`
	Job  string `json:"job" binding:"required,oneof=designer developer"`
}

func (h *HealthcheckHandlers) TestValidate(c *gin.Context) {

	data, err := validate.BindJSON[TestObj](c)
	if err != nil {
		c.JSON(http.StatusBadRequest, validate.ResponseReturn(err, data))
	}

	h.logger.Info("Paramster Success...")

	c.JSON(http.StatusOK, validate.ResponseReturn(err, data))
}

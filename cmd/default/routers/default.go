package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/dispatcher/cmd/default/handlers"
	"github.com/zkfmapf123/dispatcher/middlewares"
	"go.uber.org/zap"
)

func DefaultRouter(r *gin.Engine, prefix string, serverTimeout time.Duration, logger zap.Logger, job chan<- func()) {

	group := r.Group(prefix,
		middlewares.TimeoutMiddleware(serverTimeout))

	// handlers
	h := handlers.NewHealthCheckHandlers(logger, job)

	group.GET("", h.HealthCheck)
	group.GET("/readness", h.Readiness)
	group.GET("/liveness", h.Liveness)

	// test
	group.GET("/timeout-test", h.TimeoutTest)
	group.GET("/worker-test", h.TestWorker)
	group.POST("/validate", h.TestValidate)
}

package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/dispatcher/cmd/default/handlers"
	"github.com/zkfmapf123/dispatcher/middlewares"
	"go.uber.org/zap"
)

func DefaultRouter(r *gin.Engine, prefix string, serverTimeout time.Duration, logger zap.Logger) {

	group := r.Group(prefix)

	// handlers
	healthCheckHandlers := handlers.NewHealthCheckHandlers(logger)

	group.GET("", middlewares.TimeoutMiddleware(serverTimeout), healthCheckHandlers.HealthCheck)
	group.GET("/readness", middlewares.TimeoutMiddleware(serverTimeout), healthCheckHandlers.Readiness)
	group.GET("/liveness", middlewares.TimeoutMiddleware(serverTimeout), healthCheckHandlers.Liveness)

	// test
	group.GET("/timeout-test", middlewares.TimeoutMiddleware(serverTimeout), healthCheckHandlers.TimeoutTest)
}

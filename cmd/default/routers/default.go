package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/dispatcher/cmd/default/handlers"
	"go.uber.org/zap"
)

func DefaultRouter(r *gin.Engine, prefix string, logger zap.Logger) {

	group := r.Group(prefix)

	// handlers
	healthCheckHandlers := handlers.NewHealthCheckHandlers(logger)

	group.GET("/", healthCheckHandlers.HealthCheck)
	group.GET("/readness", healthCheckHandlers.Readiness)
	group.GET("/liveness", healthCheckHandlers.Liveness)
}

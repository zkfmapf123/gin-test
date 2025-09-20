package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TimerMiddleware(logger zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		elapsed := time.Since(start)
		logger.Info("Server Req / Res Time", zap.String("path", c.Request.URL.Path), zap.String("elapsed", elapsed.String()))
	}
}

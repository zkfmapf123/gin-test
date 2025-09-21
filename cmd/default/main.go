package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/dispatcher/cmd/default/routers"
	"github.com/zkfmapf123/dispatcher/internal/concurrency"
	"github.com/zkfmapf123/dispatcher/internal/secrets"
	"github.com/zkfmapf123/dispatcher/middlewares"
	"go.uber.org/zap"
)

var (
	// Application
	SERVER_READ_TIMEOUT  = 10 * time.Second
	SERVER_WRITE_TIMEOUT = 10 * time.Second
	SERVER_IDLE_TIMEOUT  = 60 * time.Second

	// Server
	SERVER_API_TIMEOUT = 10 * time.Second

	// Gracefully Shutdown
	SERVER_GRACE_SHUTDOWN_TIMEOUT = 30 * time.Second
	JOB_COUNT                     = 100
)

var (
	envMapping = map[string]string{
		"":    ".dev.env",
		"dev": ".dev.env",
		"stg": ".stg.env",
		"prd": ".env",
	}
)

func main() {

	// environment
	env := os.Getenv("ENV")
	secrets.SetValue(filepath.Join(envMapping[env]))

	// logger
	logger := secrets.NewLogger()
	defer logger.Sync()

	job := make(chan func(), JOB_COUNT)
	go concurrency.Dispatcher(logger, job)

	// router
	router := getRouter(logger, job)
	server := serverSetting(router)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// shutdown
	go func() {
		logger.Info("server starting", zap.String("server", env), zap.String("port", viper.GetString("PORT")))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), SERVER_GRACE_SHUTDOWN_TIMEOUT)
	defer cancel()

	close(job)

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", zap.Error(err))
	}

	logger.Info("server gracefully stopped")
}

func getRouter(logger zap.Logger, job chan<- func()) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.TimerMiddleware(logger))

	// Router Group
	routers.DefaultRouter(r, "/health", SERVER_API_TIMEOUT, logger, job)

	return r
}

func serverSetting(r *gin.Engine) *http.Server {

	return &http.Server{
		Addr:           fmt.Sprintf(":%s", viper.Get("PORT")),
		Handler:        r,
		ReadTimeout:    SERVER_READ_TIMEOUT,  // 클라이언트가 요청을 보내는 제한 시간
		WriteTimeout:   SERVER_WRITE_TIMEOUT, // 서버가 응답을 보내는 시간 제한
		IdleTimeout:    SERVER_IDLE_TIMEOUT,  // Keep-Alive 가 유지되는 시간
		MaxHeaderBytes: 1 << 20,              // x 1MB
	}

}

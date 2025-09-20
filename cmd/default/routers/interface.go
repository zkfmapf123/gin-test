package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterParams struct {
	Prefix        string
	r             *gin.Engine
	logger        *zap.Logger
	serverTimeout int
	job           chan func()
}

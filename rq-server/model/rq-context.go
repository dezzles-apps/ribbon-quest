package model

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RQContext struct {
	Logger *zap.Logger
	Ctx    *context.Context
}

func GetContext(c *gin.Context) *RQContext {
	var logger *zap.Logger
	loggerVal, exists := c.Get("logger")
	if exists {
		logger = loggerVal.(*zap.Logger)
	} else {
		logger = zap.NewNop()
	}
	ctx := c.Request.Context()
	return &RQContext{
		Logger: logger,
		Ctx:    &ctx,
	}
}

package interfaces

import (
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services"

	"github.com/gin-gonic/gin"
)

type ControllerInitContext interface {
	Router() *gin.Engine
	AuthMiddleware() *middleware.AuthMiddleware
	Config() *model.AppConfig
	Services() *services.ServiceContext
}

type Controller interface {
	Initialise(ControllerInitContext)
}

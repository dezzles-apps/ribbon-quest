package initialisers

import (
	"github.com/dezzles-apps/go-common/db"
	"github.com/gin-gonic/gin"

	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/repositories"
	"dezzles-apps/rq-server/services"
)

type InitContext struct {
	Router         *gin.Engine
	AuthMiddleware *middleware.AuthMiddleware
	Database       *db.Database

	Repositories *repositories.RepositoryContext
	Services     *services.ServiceContext
}

func (ctx *InitContext) Initialise(
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	database *db.Database,
) {
	ctx.Router = router
	ctx.AuthMiddleware = authMiddleware
	ctx.Database = database
	ctx.Repositories = repositories.InitaliseRepositories(ctx.Database)
	ctx.Services = &services.ServiceContext{}
	ctx.Services.Initialise(ctx.Database, *ctx.Repositories)

}

package controllers

import (
	"dezzles-apps/rq-server/interfaces"
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services"

	"dezzles-apps/rq-server/controllers/data"
	"dezzles-apps/rq-server/controllers/pokedexes"
	"dezzles-apps/rq-server/controllers/ribbons"
	"dezzles-apps/rq-server/controllers/shinies"

	"github.com/gin-gonic/gin"
)

type InitContext struct {
	router         *gin.Engine
	authMiddleware *middleware.AuthMiddleware
	config         *model.AppConfig
	services       *services.ServiceContext
}

func InitialiseControllers(
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	config *model.AppConfig,
	services *services.ServiceContext,
) {
	ctx := &InitContext{
		router:         router,
		authMiddleware: authMiddleware,
		config:         config,
		services:       services,
	}
	controllers := []interfaces.Controller{
		&data.PokemonController{},
		&pokedexes.PokedexController{},
		&ribbons.GamesController{},
		&ribbons.PokemonController{},
		&ribbons.RibbonStatsController{},
		&shinies.ShinyController{},
		&AuthController{},
		&EventsController{},
	}

	for _, controller := range controllers {
		controller.Initialise(ctx)
	}

}

func (ctx *InitContext) AuthMiddleware() *middleware.AuthMiddleware {
	return ctx.authMiddleware
}

func (ctx *InitContext) Config() *model.AppConfig {
	return ctx.config
}

func (ctx *InitContext) Router() *gin.Engine {
	return ctx.router
}

func (ctx *InitContext) Services() *services.ServiceContext {
	return ctx.services
}

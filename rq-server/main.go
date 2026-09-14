package main

import (
	"dezzles-apps/rq-server/controllers"
	"dezzles-apps/rq-server/initialisers"
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services"

	"log"

	"github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
)

var database *db.Database = &db.Database{}
var configService *services.ConfigService = services.NewConfigService(database)
var userService *services.UserService = services.NewUserService(database, configService)
var eventService *services.EventService = services.NewEventService(database)

func main() {
	config, err := cmodel.LoadConfig[model.AppConfig]()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	err = database.Connect(&config.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	authMiddleware := middleware.NewAuthMiddleware(&config.App)

	router := gin.Default()
	router.Use(ErrorHandler())

	initialisers.InitialiseRibbons(router, authMiddleware, database)

	controllers.NewAuthController(router, &config.App, userService)
	controllers.NewEventsController(router, *eventService)
	router.Run(":8083")
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{"errors": c.Errors})
		}
	}
}

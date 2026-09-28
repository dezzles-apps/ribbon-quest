package main

import (
	"dezzles-apps/rq-server/controllers"
	"dezzles-apps/rq-server/initialisers"
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services"
	"os"

	"log"

	"github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/dezzles-apps/go-common/monitoring"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
)

var monitor = monitoring.Monitoring{}
var database *db.Database = &db.Database{}
var configService *services.ConfigService = services.NewConfigService(database)
var userService *services.UserService = services.NewUserService(database, configService)
var eventService *services.EventService = services.NewEventService(database)

func newResource() *resource.Resource {
	hostName, _ := os.Hostname()
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceVersion("1.0.0"),
		semconv.HostName(hostName),
	)
}

func main() {
	monitor.Initialise()
	defer monitor.Shutdown()
	z, _ := zap.NewProduction()
	monitor.Logger.Info("Starting app")
	monitor.Logger = z
	config, err := cmodel.LoadConfig[model.AppConfig]()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	err = database.Connect(&config.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	monitor.Logger.Info("Database connected")
	authMiddleware := middleware.NewAuthMiddleware(&config.App)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	monitor.ConfigureGinRouter(router)
	router.Use(ErrorHandler())

	initialisers.InitialiseRibbons(router, authMiddleware, database)
	initialisers.InitialiseData(router, database)

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

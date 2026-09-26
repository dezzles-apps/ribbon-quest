package main

import (
	"context"
	"dezzles-apps/rq-server/controllers"
	"dezzles-apps/rq-server/initialisers"
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services"
	"os"

	"log"

	"github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
	"github.com/hyperdxio/opentelemetry-go/otelzap"
	"github.com/hyperdxio/opentelemetry-logs-go/exporters/otlp/otlplogs"
	sdk "github.com/hyperdxio/opentelemetry-logs-go/sdk/logs"
	"github.com/hyperdxio/otel-config-go/otelconfig"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
)

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
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("Error configuring otel: %s", err.Error())
	}
	defer otelShutdown()

	ctx := context.Background()
	logExporter, err := otlplogs.NewExporter(ctx)
	if err != nil {
		log.Fatalf("Error configuring OTLP log exporter: %s", err.Error())
	}
	loggerProvider := sdk.NewLoggerProvider(
		sdk.WithBatcher(logExporter),
	)
	defer loggerProvider.Shutdown(ctx)

	logger := zap.New(otelzap.NewOtelCore(loggerProvider))
	zap.ReplaceGlobals(logger)

	logger.Info("Starting app")
	config, err := cmodel.LoadConfig[model.AppConfig]()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	err = database.Connect(&config.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	logger.Info("Database connected")
	authMiddleware := middleware.NewAuthMiddleware(&config.App)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME")))
	router.Use(middleware.WithTraceMetadata(logger))
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

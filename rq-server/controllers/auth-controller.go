package controllers

import (
	"dezzles-apps/rq-server/interfaces"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/services"

	"net/http"
	"time"

	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct {
	config      *model.AppConfig
	userService *services.UserService
}

func (ac *AuthController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	var controller = &AuthController{
		config:      ctx.Config(),
		userService: ctx.Services().Users,
	}
	controller.registerRoutes(ctx)
}

func (ac *AuthController) registerRoutes(ctx interfaces.ControllerInitContext) {
	authGroup := ctx.Router().Group("/api/auth/v1")
	{
		authGroup.POST("/register", ac.register)
		authGroup.POST("/login", ac.login)
	}
}

func (ac *AuthController) register(c *gin.Context) {
	var input model.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": errors.InvalidRequestBody})
		return
	}
	ctx := cmodel.GetContext(c)
	err := ac.userService.RegisterUser(ctx, input)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"errors": err.Errors()})
		return
	}
	c.JSON(200, gin.H{"message": "user registered successfully"})
}

func (ac *AuthController) login(c *gin.Context) {
	var input model.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": errors.InvalidRequestBody})
		return
	}
	ctx := cmodel.GetContext(c)
	err := ac.userService.LoginUser(ctx, input)
	if err != nil {
		c.JSON(401, gin.H{"errors": err.Errors()})
		return
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  input.Username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	token, e := generateToken.SignedString([]byte(ac.config.JwtSecret))

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"errors": errors.InternalServerError.Errors()}})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

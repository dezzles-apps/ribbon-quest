package controllers

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct {
	config      *model.AppConfig
	userService *services.UserService
}

func NewAuthController(
	router *gin.Engine,
	config *model.AppConfig,
	userService *services.UserService,
) *AuthController {
	var controller = &AuthController{
		config:      config,
		userService: userService,
	}
	controller.registerRoutes(router)
	return controller
}

func (ac *AuthController) registerRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth/v1")
	{
		authGroup.POST("/register", ac.register)
		authGroup.POST("/login", ac.login)
	}
}

func (ac *AuthController) register(c *gin.Context) {
	var input model.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := ac.userService.RegisterUser(input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user registered successfully"})
}

func (ac *AuthController) login(c *gin.Context) {
	var input model.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := ac.userService.LoginUser(input)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  input.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(ac.config.JwtSecret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

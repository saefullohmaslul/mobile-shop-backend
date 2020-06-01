package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db/entity"
	"github.com/saefullohmaslul/mobile-shop-backend/src/services"
)

// AuthController is controller to handle authentication
type AuthController struct {
	Service services.AuthService
}

// NewAuthController is constructor to create auth controller instance
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		Service: *authService,
	}
}

// Register is controller to handle registration proccess
func (ctl *AuthController) Register(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	token := ctl.Service.Register(user)

	response.Success(c, "Success register user", token)
}

// Login is controller to handle login process
func (ctl *AuthController) Login(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	token := ctl.Service.Login(user)

	response.Success(c, "Success login user", token)
}

// RefreshToken is controller to handle generated access token
func (ctl *AuthController) RefreshToken(c *gin.Context) {
	var authInformation entity.AuthInformation
	_ = c.ShouldBindBodyWith(&authInformation, binding.JSON)

	token := ctl.Service.RefreshToken(authInformation)

	response.Success(c, "Success generate access token", token)
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController is controller to handle authentication
type AuthController struct {
}

// NewAuthController is constructor to create auth controller instance
func NewAuthController() *AuthController {
	return &AuthController{}
}

// Register is controller to handle registration proccess
func (ctl *AuthController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success register user",
		"data":    "Success",
		"errors":  nil,
	})
}

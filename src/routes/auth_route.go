package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/mobile-shop-backend/src/controllers"
)

// AuthRouter is method to initialize auth route
func AuthRouter(g *gin.RouterGroup) {
	authController := controllers.NewAuthController()
	{
		g.POST("/auth/register", authController.Register)
	}
}

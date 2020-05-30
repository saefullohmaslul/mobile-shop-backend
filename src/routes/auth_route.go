package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/mobile-shop-backend/src/controllers"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db"
	"github.com/saefullohmaslul/mobile-shop-backend/src/repositories"
	"github.com/saefullohmaslul/mobile-shop-backend/src/services"
)

// AuthRouter is method to initialize auth route
func AuthRouter(g *gin.RouterGroup) {
	conn := db.DB{}
	userRepository := repositories.NewUserRepository(conn.Get())
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)
	{
		g.POST("/auth/register", authController.Register)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/mobile-shop-backend/src/controllers"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db"
	"github.com/saefullohmaslul/mobile-shop-backend/src/repositories"
	"github.com/saefullohmaslul/mobile-shop-backend/src/services"
	"github.com/saefullohmaslul/mobile-shop-backend/src/validations"
)

// AuthRouter is method to initialize auth route
func AuthRouter(g *gin.RouterGroup) {
	conn := db.DB{}
	userRepository := repositories.NewUserRepository(conn.Get())
	authInformationRepository := repositories.NewAuthInformationRepository(conn.Get())
	authService := services.NewAuthService(userRepository, authInformationRepository)
	authController := controllers.NewAuthController(authService)
	{
		g.POST("/auth/register", validations.Register, authController.Register)
		g.POST("/auth/login", validations.Login, authController.Login)
	}
}

package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/mobile-shop-backend/src/routes"
)

// Application is application struct for init server
type Application struct {
	Route *gin.Engine
}

// NewApplication is constructor to create application
func NewApplication(route *gin.Engine) *Application {
	return &Application{Route: route}
}

// Create is method to create server application
func (a Application) Create() {
	configureEndpoint(a.Route)
}

func configureEndpoint(r *gin.Engine) {
	g := r.Group("/api/v1")

	routes.Router(g)
	routes.NoRoute(r)
}

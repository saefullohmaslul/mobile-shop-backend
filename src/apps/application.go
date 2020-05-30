package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/jpoles1/gopherbadger/logging"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db"
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
	a.Route.Use(response.Recovery(response.ErrorHandler))
	configureDB()
	configureEndpoint(a.Route)
}

func configureEndpoint(r *gin.Engine) {
	g := r.Group("/api/v1")

	routes.HealthRouter(g)
	routes.AuthRouter(g)
	routes.NoRoute(r)
}

func configureDB() {
	_, err := db.NewDB()
	if err != nil {
		logging.Error("DB", err)
	}
}

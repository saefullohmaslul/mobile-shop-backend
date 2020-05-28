package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoRoute is method to handle route not found
func NoRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route not found",
			"data":    nil,
			"errors": []map[string]interface{}{
				gin.H{
					"flag":    "ROUTE_NOT_FOUND",
					"message": "The route you are looking for is not found",
				},
			},
		})
	})
}

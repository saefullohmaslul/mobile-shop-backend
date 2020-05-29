package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthRouter is method to initialize health checking
func HealthRouter(g *gin.RouterGroup) {
	{
		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Health OK",
				"data": gin.H{
					"database": "OK",
				},
				"errors": nil,
			})
		})
	}
}

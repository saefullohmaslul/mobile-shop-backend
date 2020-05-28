package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router is method to initialize router
func Router(g *gin.RouterGroup) {
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

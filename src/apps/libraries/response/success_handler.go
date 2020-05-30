package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success is success response
func Success(c *gin.Context, message string, data interface{}) {
	response := Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
		Errors:  nil,
	}

	c.JSON(http.StatusOK, response)
}

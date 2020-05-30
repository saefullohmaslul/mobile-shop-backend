package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// ErrorHandler is middleware to sending error response
func ErrorHandler(c *gin.Context, err interface{}) {
	response := Response{}
	var errors []Error

	if err := mapstructure.Decode(err, &response); err != nil {
		errors = append(errors, Error{
			Message: "An error occurred on our server", Flag: "ERROR_MAP_TO_STRUCT",
		})
		res := Response{
			Status:  http.StatusInternalServerError,
			Message: response.Message,
			Data:    nil,
			Errors:  errors,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)

		return
	}

	res := Response{
		Status:  response.Status,
		Message: response.Message,
		Data:    nil,
		Errors:  response.Errors,
	}
	c.AbortWithStatusJSON(response.Status, res)
}

package response

import (
	"github.com/gin-gonic/gin"
)

// Recovery is middleware to use custom error response
func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithoutWriter(f)
}

// RecoveryWithoutWriter is recover panic middleware
func RecoveryWithoutWriter(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				f(c, err)
			}
		}()

		/**
		 * forward to next middleware
		 */
		c.Next()
	}
}

package generator

import (
	"time"

	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/token"
)

// Token is generator for access token and refresh token
func Token(payload interface{}) (string, time.Time, string) {
	var errors []response.Error

	token, expiresIn, err := token.SignJWT(payload)
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "TOKEN_SIGN_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	refreshToken := GenerateHMAC(payload)

	return token, expiresIn, refreshToken
}

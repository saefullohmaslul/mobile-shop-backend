package validations

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/validations/schemas"
)

// Register is validation for register endpoint
func Register(c *gin.Context) {
	var errors []response.Error
	var register schemas.Register

	if err := c.ShouldBindBodyWith(&register, binding.JSON); err != nil {
		errors = append(errors, response.Error{
			Message: err.Error(),
			Flag:    "INVALID_VALIDATION_JSON",
		})
		response.BadRequest("Validation error", errors)
	}

	registerValidate := &schemas.Register{
		Name:     register.Name,
		UserName: register.UserName,
		Email:    register.Email,
		Password: register.Password,
	}

	Validate(registerValidate, errors)
}

// Login is validation for login endpoint
func Login(c *gin.Context) {
	var errors []response.Error
	var login schemas.Login

	if err := c.ShouldBindBodyWith(&login, binding.JSON); err != nil {
		errors = append(errors, response.Error{
			Message: err.Error(),
			Flag:    "INVALID_VALIDATION_JSON",
		})
		response.BadRequest("Validation error", errors)
	}

	loginValidate := &schemas.Login{
		UserName: login.UserName,
		Password: login.Password,
	}

	Validate(loginValidate, errors)
}

// RefreshToken is validation for refresh token endpoint
func RefreshToken(c *gin.Context) {
	var errors []response.Error
	var refreshToken schemas.RefreshToken

	if err := c.ShouldBindBodyWith(&refreshToken, binding.JSON); err != nil {
		errors = append(errors, response.Error{
			Message: err.Error(),
			Flag:    "INVALID_VALIDATION_JSON",
		})
		response.BadRequest("Validation error", errors)
	}

	if refreshToken.RefreshToken == "" {
		errors = append(errors, response.Error{
			Message: "RefreshToken field required",
			Flag:    fmt.Sprintf("INVALID_VALIDATION_%s", "REFRESH_TOKEN"),
		})
	}

	isMatch, err := regexp.Match("^[A-Za-z0-9-_=]+\\.[0-9]*$", []byte(refreshToken.RefreshToken))
	if err != nil {
		errors = append(errors, response.Error{
			Message: err.Error(),
			Flag:    fmt.Sprintf("INVALID_VALIDATION_%s", "REFRESH_TOKEN"),
		})
	}

	if !isMatch {
		errors = append(errors, response.Error{
			Message: "Refresh token must be refresh token pattern",
			Flag:    fmt.Sprintf("INVALID_VALIDATION_%s", "REFRESH_TOKEN"),
		})
	}

	if errors != nil {
		response.BadRequest("Validation error", errors)
	}
}

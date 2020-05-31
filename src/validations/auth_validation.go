package validations

import (
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

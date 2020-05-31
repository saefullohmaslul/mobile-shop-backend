package services

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/token"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db/entity"
	"github.com/saefullohmaslul/mobile-shop-backend/src/helpers/generator"
	"github.com/saefullohmaslul/mobile-shop-backend/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is service to handle authentication flow
type AuthService struct {
	UserRepository repositories.UserRepository
}

// AuthReturn is return format for authentication token
type AuthReturn struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    time.Time `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
}

// NewAuthService is constructor to create auth service instance
func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: *userRepository,
	}
}

// Register is service to handle logic of registration
func (s *AuthService) Register(user entity.User) *AuthReturn {
	var errors []response.Error

	userExist := checkUserExist(s.UserRepository, user)
	if (userExist != entity.User{}) {
		errors = append(errors, response.Error{
			Flag:    "USER_ALREADY_EXIST",
			Message: "User with this username already exist",
		})
		response.Conflict("User already exist", errors)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "HASH_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	if err := s.UserRepository.Register(entity.User{
		Username: user.Username,
		Password: string(hash),
		Name:     user.Name,
	}); err != nil {
		errors = append(errors, response.Error{
			Flag:    "USER_REGISTER_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	return generateToken(user)
}

// Login is service to handle logic of login
func (s *AuthService) Login(user entity.User) *AuthReturn {
	var errors []response.Error

	userExist := checkUserExist(s.UserRepository, user)
	if (userExist == entity.User{}) {
		errors = append(errors, response.Error{
			Flag:    "USER_NOT_FOUND",
			Message: "User with this username not found",
		})
		response.NotFound("User not found", errors)
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(user.Password))
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "USER_PASSWORD_NOT_MATCH",
			Message: "Password wrong",
		})
		response.Unauthorized("Password not match", errors)
	}

	return generateToken(user)
}

func checkUserExist(repository repositories.UserRepository, user entity.User) entity.User {
	var errors []response.Error

	userExist, err := repository.UserExist(entity.User{Username: user.Username})
	if gorm.IsRecordNotFoundError(err) {
		return userExist
	}
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "USER_CHECK_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}
	return userExist
}

func generateToken(user entity.User) *AuthReturn {
	var errors []response.Error

	payload := map[string]interface{}{"username": user.Username}
	token, expiresIn, err := token.SignJWT(payload)
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "TOKEN_SIGN_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	refreshToken := generator.GenerateHMAC(payload)

	return &AuthReturn{
		AccessToken:  token,
		TokenType:    "bearer",
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
	}
}

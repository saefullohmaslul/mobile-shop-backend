package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db/entity"
	"github.com/saefullohmaslul/mobile-shop-backend/src/helpers/generator"
	"github.com/saefullohmaslul/mobile-shop-backend/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is service to handle authentication flow
type AuthService struct {
	UserRepository            repositories.UserRepository
	AuthInformationRepository repositories.AuthInformationRepository
}

// AuthReturn is return format for authentication token
type AuthReturn struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    time.Time `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
}

// RefreshTokenReturn is return format for refresh token endpoint
type RefreshTokenReturn struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   time.Time `json:"expires_in"`
}

// NewAuthService is constructor to create auth service instance
func NewAuthService(userRepository *repositories.UserRepository, authInformationRepository *repositories.AuthInformationRepository) *AuthService {
	return &AuthService{
		UserRepository:            *userRepository,
		AuthInformationRepository: *authInformationRepository,
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
			Flag:    "BCRYPT_HASH_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	userCreated, err := s.UserRepository.Register(entity.User{
		Username: user.Username,
		Password: string(hash),
		Name:     user.Name,
	})
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "USER_REGISTER_DB_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	payload := map[string]interface{}{"username": user.Username}
	accessToken, expiresIn := generator.Token(payload)
	refreshToken := generator.GenerateHMAC(payload)

	generator.StoreRefreshToken(userCreated.ID, s.AuthInformationRepository, refreshToken, errors)

	return &AuthReturn{
		AccessToken:  accessToken,
		TokenType:    "bearer",
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
	}
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

	payload := map[string]interface{}{"username": user.Username}
	accessToken, expiresIn := generator.Token(payload)
	refreshToken := generator.GenerateHMAC(payload)

	generator.StoreRefreshToken(userExist.ID, s.AuthInformationRepository, refreshToken, errors)

	return &AuthReturn{
		AccessToken:  accessToken,
		TokenType:    "bearer",
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
	}
}

// RefreshToken is
func (s *AuthService) RefreshToken(authInformation entity.AuthInformation) *RefreshTokenReturn {
	var errors []response.Error

	refreshToken := strings.Split(authInformation.RefreshToken, ".")
	expiredTime, _ := strconv.ParseInt(refreshToken[1], 10, 64)

	if expiredTime < time.Now().Unix() {
		errors = append(errors, response.Error{
			Flag:    "REFRESH_TOKEN_EXPIRED",
			Message: "Refresh token expired, please login again",
		})
		response.Unauthorized("Refresh token expired", errors)
	}

	userAuthInformation, err := s.AuthInformationRepository.GetUser(authInformation)
	if gorm.IsRecordNotFoundError(err) {
		errors = append(errors, response.Error{
			Flag:    "USER_NOT_FOUND",
			Message: "User with this refresh token not found",
		})
		response.NotFound("User not found", errors)
	}

	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "REFRESH_TOKEN_DB_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	payload := map[string]interface{}{"username": userAuthInformation.User.Username}
	accessToken, expiresIn := generator.Token(payload)

	return &RefreshTokenReturn{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
		TokenType:   "bearer",
	}
}

func checkUserExist(repository repositories.UserRepository, user entity.User) entity.User {
	var errors []response.Error

	userExist, err := repository.UserExist(entity.User{Username: user.Username})
	if gorm.IsRecordNotFoundError(err) {
		return userExist
	}
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "USER_CHECK_DB_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}
	return userExist
}

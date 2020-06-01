package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db/entity"
)

// AuthInformationRepository is layer to execute sql for table auth_informations
type AuthInformationRepository struct {
	Conn *gorm.DB
}

// NewAuthInformationRepository is constructor to create auth information instance
func NewAuthInformationRepository(conn *gorm.DB) *AuthInformationRepository {
	return &AuthInformationRepository{
		Conn: conn,
	}
}

// CreateRefreshToken is method to create refresh token
func (r *AuthInformationRepository) CreateRefreshToken(authInformation entity.AuthInformation) error {
	if err := r.Conn.Create(&authInformation).Error; err != nil {
		return err
	}
	return nil
}

// UpdateRefreshToken is method to create refresh token
func (r *AuthInformationRepository) UpdateRefreshToken(param, authInformation entity.AuthInformation) error {
	if err := r.Conn.Model(&param).Update(&authInformation).Error; err != nil {
		return err
	}
	return nil
}

// GetUserID is method to get user_id with refresh token
func (r *AuthInformationRepository) GetUserID(param entity.AuthInformation) (entity.AuthInformation, error) {
	authInformation := entity.AuthInformation{}
	if err := r.Conn.Select("user_id").Where(&param).First(&authInformation).Error; err != nil {
		return authInformation, err
	}
	return authInformation, nil
}

// GetUser is method to get user_id with refresh token
func (r *AuthInformationRepository) GetUser(param entity.AuthInformation) (entity.AuthInformation, error) {
	authInformation := entity.AuthInformation{}
	if err := r.Conn.Preload("User").
		Where("refresh_token = ?", param.RefreshToken).
		First(&authInformation).
		Error; err != nil {
		return authInformation, err
	}
	return authInformation, nil
}

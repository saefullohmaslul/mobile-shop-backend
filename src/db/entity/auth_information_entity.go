package entity

import uuid "github.com/satori/go.uuid"

// AuthInformation is struct for auth_informations entity
type AuthInformation struct {
	Base
	UserID       uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
	RefreshToken string    `gorm:"type:varchar(255)" json:"refresh_token"`
	User         User
}

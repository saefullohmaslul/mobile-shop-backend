package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db/entity"
)

// UserRepository is layer to execution sql for table users
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository is constructor to create user repository instance
func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

// Register is method to create user in db
func (r *UserRepository) Register(user entity.User) (entity.User, error) {
	userCreated := entity.User{}
	if err := r.Conn.Create(&user).Scan(&userCreated).Error; err != nil {
		return userCreated, err
	}

	return userCreated, nil
}

// UserExist is method to check existing user by username or email
func (r *UserRepository) UserExist(param entity.User) (entity.User, error) {
	user := entity.User{}
	var err error

	switch param.Email {
	case nil:
		err = r.Conn.Select("id, username, password").Where(&entity.User{Username: param.Username}).First(&user).Error
	default:
		err = r.Conn.Select("id, email, password").Where(&entity.User{Email: param.Email}).First(&user).Error
	}

	if err != nil {
		return user, err
	}
	return user, nil
}

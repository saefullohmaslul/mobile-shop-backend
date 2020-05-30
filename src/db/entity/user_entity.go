package entity

// User is struct for user entity
type User struct {
	Base
	Username *string `gorm:"type:varchar(50);unique_index" json:"username"`
	Name     *string `json:"name" faker:"name"`
	Email    *string `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Password *string `gorm:"type:varchar(255)" json:"password" faker:"password"`
}

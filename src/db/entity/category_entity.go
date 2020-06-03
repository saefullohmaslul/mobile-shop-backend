package entity

// Category is struct for category entity
type Category struct {
	Base
	Name string `gorm:"type:varchar(50)" json:"name"`
}

package entity

import uuid "github.com/satori/go.uuid"

// Cart is struct for cart entity
type Cart struct {
	Base
	UserID    uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid;unique" json:"product_id"`
	Amount    int16     `gorm:"type:integer" json:"amount"`
	SubTotal  float64   `gorm:"type:float" json:"sub_total"`
}

package entity

import uuid "github.com/satori/go.uuid"

// CartProductEntity is struct for cart product entity
type CartProductEntity struct {
	Base
	CartID    uuid.UUID `gorm:"type:uuid;unique" json:"cart_id"`
	ProductID uuid.UUID `gorm:"type:uuid;unique" json:"product_id"`
}

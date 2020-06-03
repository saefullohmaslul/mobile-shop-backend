package entity

import uuid "github.com/satori/go.uuid"

// ProductLiked is struct for product entity
type ProductLiked struct {
	Base
	ProductID uuid.UUID `gorm:"type:uuid;unique" json:"product_id"`
	UserID    uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
}

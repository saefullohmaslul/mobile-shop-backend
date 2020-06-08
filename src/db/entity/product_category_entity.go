package entity

import uuid "github.com/satori/go.uuid"

// ProductCategory is struct for product category entity
type ProductCategory struct {
	Base
	ProductID  uuid.UUID `gorm:"type:uuid;unique" json:"product_id"`
	CategoryID uuid.UUID `gorm:"type:uuid;unique" json:"category_id"`
}

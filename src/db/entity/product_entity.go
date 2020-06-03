package entity

import uuid "github.com/satori/go.uuid"

// Product is struct for product entity
type Product struct {
	Base
	Thumbnail         string    `gorm:"type:text" json:"thumbnail"`
	Title             string    `gorm:"type:varchar(50)" json:"title"`
	Description       string    `gorm:"type:text" json:"description"`
	Price             float64   `gorm:"float" json:"price"`
	ProductCategoryID uuid.UUID `gorm:"type:uuid;unique" json:"product_category_id"`
	Category          Category
}

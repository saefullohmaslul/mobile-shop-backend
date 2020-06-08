package entity

import uuid "github.com/satori/go.uuid"

// Order is struct for order entity
type Order struct {
	Base
	UserID            uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
	ProductID         uuid.UUID `gorm:"type:uuid;unique" json:"product_id"`
	Thumbnail         string    `gorm:"type:text" json:"thumbnail"`
	Title             string    `gorm:"type:varchar(50)" json:"title"`
	Description       string    `gorm:"type:text" json:"description"`
	Price             float64   `gorm:"float" json:"price"`
	ProductCategoryID uuid.UUID `gorm:"type:uuid;unique" json:"product_category_id"`
	Amount            int16     `gorm:"type:integer" json:"amount"`
	SubTotal          float64   `gorm:"type:float" json:"sub_total"`
	TransactionID     uuid.UUID `gorm:"type:uuid;unique" json:"transaction_id"`
}

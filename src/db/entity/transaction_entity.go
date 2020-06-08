package entity

import uuid "github.com/satori/go.uuid"

// Transaction is struct for transaction entity
type Transaction struct {
	Base
	UserID uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
	Total  float64   `gorm:"type:float" json:"total"`
	Status string    `gorm:"type:varchar(10)" json:"status"`
}

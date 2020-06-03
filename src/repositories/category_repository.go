package repositories

import "github.com/jinzhu/gorm"

// CategoryRepository is layer to execute sql for table categories
type CategoryRepository struct {
	Conn *gorm.DB
}

// NewCategoryRepository is constructor to create category instance
func NewCategoryRepository(conn *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		Conn: conn,
	}
}

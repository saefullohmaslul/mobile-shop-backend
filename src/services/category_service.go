package services

import "github.com/saefullohmaslul/mobile-shop-backend/src/repositories"

// CategoryService is logic for category
type CategoryService struct {
	CategoryRepository repositories.CategoryRepository
}

// NewCategoryService is constructor to create category service instance
func NewCategoryService(categoryService *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: *categoryService,
	}
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/mobile-shop-backend/src/services"
)

// CategoryController is controller to handle category request-response
type CategoryController struct {
	Service services.CategoryService
}

// NewCategoryController is constructor to create category instance
func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		Service: *categoryService,
	}
}

// CreateCategory is controller to handle category creations
func (ctl *CategoryController) CreateCategory(c *gin.Context) {

}

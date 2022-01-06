package service

import (
	"github.com/Cxz-web/go-micro2/user/domain/catefory/domain/model"
	"github.com/Cxz-web/go-micro2/user/domain/catefory/domain/repository"
)

type CategoryService struct {
	repository *repository.CategoryRepository
}

func (c *CategoryService) FindCategoryById(id int64) (category *model.Category, err error) {
	return c.repository.FindCategoryById(id)
}

func (c *CategoryService) CreateCategory(category *model.Category) error {
	return c.repository.CreateCategory(category)
}

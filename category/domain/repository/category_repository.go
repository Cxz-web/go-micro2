package repository

import "github.com/jinzhu/gorm"
import "github.com/Cxz-web/go-micro2/user/domain/catefory/domain/model"

type CategoryRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (c *CategoryRepository) InitTable() error {
	c.mysqlDB.CreateTable(&model.Category{})
	return nil
}

// 寻找商品
func (c *CategoryRepository) FindCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, c.mysqlDB.First(category, id).Error
}

func (c *CategoryRepository) CreateCategory(product *model.Category) error {
	return c.mysqlDB.Create(product).Error
}

package model

type Category struct {
	// 主键
	ID int64 `gorm: "primary_key;not_null;auto_increment"`
	// 用户名称
	CategoryName string `gorm: "unique_index;not_null" json: category_name`
	// 密码
	CategoryCode string `gorm: "unique_index;not_null" json: category_code`
}

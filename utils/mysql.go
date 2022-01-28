package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func initMysql() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:watch8562871@(127.0.0.1:3306)/category?charset=utf8mb4&parseTime=True&loc=Local")

	db.SingularTable(true)

	return db, err

}

package persistence

import (
	"gin-ddd-demo/module/command/user/adapter/persistence/sqldb/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSqlDB() *gorm.DB {
	dsn := "root:000000@tcp(192.168.1.200:3306)/ddd_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.UserModel{})
	return db
}

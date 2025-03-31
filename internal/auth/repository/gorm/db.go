package gorm

import (
	"chat/internal/auth/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

// 之后再用命令行的模式，现在这个函数暂时废弃
func getDSN() string {
	return fmt.Sprintf("lhm:HcZaz3#f*uSZYCNeJ5Eu@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
}

func InitMysqlDB() error {
	db, err := gorm.Open(mysql.Open("lhm:HcZaz3#f*uSZYCNeJ5Eu@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	db.AutoMigrate(&model.User{})
	MysqlDB = db
	return nil
}

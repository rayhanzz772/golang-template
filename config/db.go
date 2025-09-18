package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/godata?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
func InitDB() error {
	return initDB()
}	
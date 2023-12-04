package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlProxy struct {
	gorm.Model
}

var SqlProxy *gorm.DB

func Init() error {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	SqlProxy = db
	return nil
}

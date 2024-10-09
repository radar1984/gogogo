package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	
	// Auto migrate models
	db.AutoMigrate(&User{}, &Character{}, &Task{}, &Tag{})
	
	return db, nil
}
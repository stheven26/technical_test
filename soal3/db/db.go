package db

import (
	"github.com/stheven26/technical_test/config"
	"github.com/stheven26/technical_test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDB() {
	config := config.Config()

	connectionString := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@(" + config.DB_HOST + ")/" + config.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.MessageOnGroup{})
}

func GetConnection() *gorm.DB {
	return db
}

package database

import (
	"github.com/chris92vr/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/companydb"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}

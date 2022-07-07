package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:rootroot@/companydb"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myungsworld/database/models"
)

var DB *gorm.DB


const User = "root"
const Password = ""
const LocalHost = "127.0.0.1:3306"
const DBName = "myungsworld"

func ConnectDB(){

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local",
		User,
		Password,
		LocalHost,
		DBName,
		)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	createTables()

}

func createTables() {
	tables := []interface{}{
		(*models.Test)(nil),


	}

	if err := DB.AutoMigrate(tables...); err != nil {
		fmt.Println(err)
		panic(err)
	}
}


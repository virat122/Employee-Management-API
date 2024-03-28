package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:Archaudhary@122@tcp(localhost:3306)/empnew?parseTime=true"

var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(" connectio failed :( ")

	}
	Database.AutoMigrate(&Employee{})

}

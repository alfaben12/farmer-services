package main

import (
	"./app/Config"
	"./app/Models/FarmLandModel"
	"./app/Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/zpay?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("status: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&FarmLandModel.FarmLand{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}

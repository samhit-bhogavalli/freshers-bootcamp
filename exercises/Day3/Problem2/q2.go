package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"local/exercises/Day3/Problem2/Config"
	"local/exercises/Day3/Problem2/Models"
	"local/exercises/Day3/Problem2/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}

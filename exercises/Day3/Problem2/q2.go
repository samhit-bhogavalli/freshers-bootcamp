package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"local/exercises/Day3/Problem2/Config"
	"local/exercises/Day3/Problem2/Models"
	"local/exercises/Day3/Problem2/Routes"
)

var err error

func main() {
	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Status:", err)
	}
	student := Models.Student{
		Model:     gorm.Model{ID: 2},
		FirstName: "jj",
		LastName:  "lk",
		DOB:       "2019-10-6",
		Address:   "mads",
		Marks: []Models.Marks{
			{
				Subject: "math",
				Score:   1,
			},
			{
				Subject: "phy",
				Score:   1,
			},
		},
	}
	Config.DB.AutoMigrate(&Models.Student{})
	//time.Sleep(time.Second * 2)
	Config.DB.AutoMigrate(&Models.Marks{})
	Config.DB.Create(&student)
	r := Routes.SetupRouter()
	//running
	r.Run()
}

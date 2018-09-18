package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/saseke/go-web/Config"
	"github.com/saseke/go-web/Models"
	"github.com/saseke/go-web/Routers"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var err error

func init() {
	gotenv.Load()
	log.Println(os.Getenv("DB_HOST"))
}
func main() {

	//Config.DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/poseidon?charset=utf8&parseTime=True&loc=Local")
	cfgStr := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_TABLE") + "?" + os.Getenv("DB_PARAM")
	fmt.Println(cfgStr)
	Config.DB, err = gorm.Open(os.Getenv("DB_DATABSE_TYPE"), cfgStr)

	if err != nil {
		fmt.Println(err)
	}
	defer Config.DB.Close()
	Config.DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "db_" + defaultTableName
	}
	Config.DB.AutoMigrate(&Models.Item{})

	r := Routers.SetUpRouter()

	r.Run()
}

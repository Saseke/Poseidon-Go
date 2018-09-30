package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/saseke/go-web/Config"
	"github.com/saseke/go-web/Constant"
	"github.com/saseke/go-web/Models"
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
	cfgStr := os.Getenv(Constant.DB_USER) + ":" + os.Getenv(Constant.DB_PWD) + "@tcp(" +
		os.Getenv(Constant.DB_HOST) + ":" + os.Getenv(Constant.DB_PORT) + ")/" +
		os.Getenv(Constant.DB_TABLE) + "?" + os.Getenv(Constant.DB_PARAM)
	fmt.Println(cfgStr)
	Config.DB, err = gorm.Open(os.Getenv(Constant.DB_TYPE), cfgStr)

	if err != nil {
		fmt.Println(err)
	}
	initRedis()
	defer Config.DB.Close()
	// show logs
	Config.DB.LogMode(true)
	Config.DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "db_" + defaultTableName
	}
	Config.DB.AutoMigrate(&Models.Item{})

	r := gin.Default()

	r.Run()
}

func initRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(Constant.REDIS_HOST) + ":" + os.Getenv(Constant.REDIS_PORT),
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

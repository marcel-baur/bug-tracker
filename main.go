package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	fmt.Println("Starting up server...")
	fmt.Println("Wait 5 seconds...")
	time.Sleep(5 * time.Second)
	db = launchDb()
	gin.ForceConsoleColor()
	r := gin.Default()
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	r.Run()
}

func launchDb() *gorm.DB {
	fmt.Println("Connect to mariadb")
	var globalDB *gorm.DB
	var err error

	globalDB, err = gorm.Open("mysql", "root:root@tcp(mariadb:3306)/bugtrack?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		log.Println("failed to connect database ... retry")
		log.Println(err)
	} else {
		fmt.Println("Successfully connect to database")
		// TODO: Migrations
	}
	return globalDB

}

package main

import (
	myredis "chat/internal/auth/repository/go-redis"
	mygorm "chat/internal/auth/repository/gorm"
	"fmt"
)

func Init() {
	err := mygorm.InitMysqlDB()
	if err != nil {
		fmt.Println("InitMysqlDB Failed!")
		return
	}
	err = myredis.InitRedisDB()
	if err != nil {
		fmt.Println("InitRedisDB Failed!")
		return
	}
}

func main() {
	Init()
}

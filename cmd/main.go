package main

import (
	"chat/internal/auth/repository/gorm"
	"fmt"
)

func main() {
	err := gorm.InitDB()
	if err != nil {
		fmt.Println("InitDB Failed!")
		return
	}
}

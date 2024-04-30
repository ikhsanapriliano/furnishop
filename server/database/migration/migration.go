package main

import (
	"fmt"
	"furnishop/server/database"
	"furnishop/server/model"
)

func main() {
	db := database.NewInitializer().Start()

	err := db.AutoMigrate(&model.User{}, &model.ShoppingCart{})
	if err != nil {
		panic(err)
	}
	fmt.Println("migrate success")
}

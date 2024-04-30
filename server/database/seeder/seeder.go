package main

import (
	"fmt"
	"furnishop/server/database"
	"furnishop/server/database/factory"
)

func main() {
	db := database.NewInitializer().Start()

	factory.NewUserFactory(db, 5, 2).Seed()
	factory.NewUserFactory(db, 2, 1).Seed()

	fmt.Println("seeding success")
}

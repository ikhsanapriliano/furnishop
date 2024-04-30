package main

import "furnishop/server/delivery"

func main() {
	delivery.NewServer().Run()
}

package main

import (
	"chal8/database"
	"chal8/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}

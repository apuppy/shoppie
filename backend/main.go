package main

import (
	"shoppie/model"
	"shoppie/router"
)

func main() {
	// DB migration
	model.AutoMigration()

	router.Start()
}

package main

import (
	"shoppie/api"
	"shoppie/model"
)

func main() {
	// DB migration
	model.AutoMigration()

	// start http server
	bindAddr := ":8080"
	api.Listen(bindAddr)
}

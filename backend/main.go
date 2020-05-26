package main

import (
	"shoppie/api"
)

func main() {
	// DB migration
	// model.AutoMigration()

	// start http server
	bindAddr := ":8080"
	api.Listen(bindAddr)
}

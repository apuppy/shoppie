package main

import (
	"shoppie/model"
)

func main() {
	// DB migration
	model.AutoMigration()

	// router.Start()
}

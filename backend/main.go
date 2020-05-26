package main

import "shoppie/api"

func main() {
	// DB migration
	// model.AutoMigration()

	// send email
	// util.SendMail()

	// start http server
	bindAddr := ":8080"
	api.Listen(bindAddr)
}

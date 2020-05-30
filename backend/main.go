package main

import (
	"shoppie/api"
	"shoppie/config"
	"strconv"
)

func main() {
	// DB migration
	// model.AutoMigration()

	// send email
	// util.SendMail()

	// read config
	config := config.Read()

	// redis cache
	// err := cache.Ping()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// start http server
	bindAddr := ":" + strconv.Itoa(config.Server.Port)
	api.Listen(bindAddr)
}

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

	// redis cache
	// err := cache.Ping()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// start http server
	bindAddr := ":" + strconv.Itoa(config.Conf.Server.Port)
	api.Listen(bindAddr)
}

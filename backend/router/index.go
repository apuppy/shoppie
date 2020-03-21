package router

import (
	"shoppie/user"

	"github.com/gin-gonic/gin"
)

// Start parse http route and launch gin framework
func Start() {

	r := gin.Default()

	r.GET("/user", user.Home)

	r.Run()
}

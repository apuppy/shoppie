package router

import (
	"net/http"
	"shoppie/user"

	"github.com/gin-gonic/gin"
)

// Start parse http route and launch gin framework
func Start() {

	r := gin.Default()

	r.GET("/", siteHome)
	r.GET("/user", user.Home)

	r.Run()
}

func siteHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to shoppie.",
	})
}

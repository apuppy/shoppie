package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User user router
func User() {
	r := gin.Default()

	r.GET("/user/", userHome)

	r.Run()
}

func userHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user router hit.",
	})
}

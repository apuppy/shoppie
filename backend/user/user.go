package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserList() {
	r := gin.Default()
	r.GET("/user/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

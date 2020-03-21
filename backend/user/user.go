package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home user index route match "/user"
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user router hit.",
	})
}

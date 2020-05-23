package router

import (
	"shoppie/api/controller"

	"github.com/gin-gonic/gin"
)

// HandlerRouter handler the routes
func HandlerRouter(bindAddr string) {
	r := gin.Default()

	r.GET("/", controller.SiteHome)
	r.POST("/upload", controller.Upload)

	// /* staff route */
	r.POST("/staff", controller.PostStaff)
	r.GET("/staff", controller.GetStaffs)
	// r.GET("/staff/{id}", controller.GetStaffs)
	// r.PUT("/staff/{id}", controller.PutStaff)
	// r.DELETE("/staff/{id}", controller.DeleteStaff)

	r.Run(bindAddr)
}

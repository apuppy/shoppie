package router

import (
	"shoppie/api/controller"

	"github.com/gin-contrib/cors"
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
	r.GET("/staff/:id", controller.GetStaff)
	r.PUT("/staff/:id", controller.PutStaff)
	r.DELETE("/staff/:id", controller.DeleteStaff)
	r.PUT("/staff-auth/:id", controller.SetPasswd)
	r.POST("/staff-auth", controller.VerifyPasswd)

	r.Use(cors.Default())
	r.Run(bindAddr)
}

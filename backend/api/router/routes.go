package router

import (
	"shoppie/api/controller"
	"shoppie/util"

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
	r.POST("/staff-login", controller.VerifyPasswd)
	r.GET("/staff-info", controller.Info)
	r.POST("/staff-logout", controller.Logout)
	// demo route
	r.GET("/demo/index", controller.Index)

	r.Use(util.GinCors())

	r.Run(bindAddr)
}

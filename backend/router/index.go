package router

import (
	"net/http"
	"shoppie/user"
	"shoppie/util"

	"github.com/gin-gonic/gin"
)

// Start parse http route and launch gin framework
func Start() {

	r := gin.Default()

	r.GET("/", siteHome)
	r.GET("/user", user.Home)
	r.POST("/upload", upload)

	r.Run()
}

func siteHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to shoppie.",
	})
}

func upload(c *gin.Context) {
	// curl --location --request POST 'http://localhost:8080/upload' \
	// --header 'Content-Type: multipart/form-data' \
	// --form 'file=@/Users/hongde/test.zip'

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

	// Upload the file to specific dst.
	dstFilename := util.GetFilename(file.Filename)
	dstExt := util.GetExt(file.Filename)
	randomString := util.RandomString(16)
	dst := "/Users/hongde/code/test/files/" + dstFilename + "-" + randomString + "." + dstExt
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"filename": file.Filename,
	})
}

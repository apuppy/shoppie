package controller

import (
	"net/http"
	"shoppie/util"

	"github.com/gin-gonic/gin"
)

// SiteHome default route
func SiteHome(c *gin.Context) {
	util.JSON(c, gin.H{"message": "Welcome to shoppie."})
}

// Upload upload file
func Upload(c *gin.Context) {
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
	dst := "/Users/hongde/code/test/files/" + dstFilename + "-" + randomString + dstExt
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

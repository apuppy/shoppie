package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// BodyParser parse http body
func BodyParser(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

// BodyParamsParser parse http body
func BodyParamsParser(c *gin.Context) []byte {
	body, _ := ioutil.ReadAll(c.Request.Body)
	return body
}

// ToJSON http json header
func ToJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json;charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckError(err)
}

// JSON gin json response
func JSON(c *gin.Context, response gin.H) {
	c.JSON(http.StatusOK, response)
}

// Success HTTP api success response
func Success(c *gin.Context, data interface{}) {
	GinCorsHeaders(c)
	response := gin.H{"code": 20000, "message": "success", "data": data}
	c.JSON(http.StatusOK, response)
}

// Error HTTP api error response
func Error(c *gin.Context, message string) {
	GinCorsHeaders(c)
	response := gin.H{"code": 40000, "message": message, "data": nil}
	c.JSON(http.StatusOK, response)
}

// CheckError catch error and return
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GinCorsHeaders add HTTP CORS headers
func GinCorsHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "x-token")
}

// GinCors gin CORS settings
func GinCors() gin.HandlerFunc {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AddAllowHeaders("x-token")
	return cors.New(corsConf)
}

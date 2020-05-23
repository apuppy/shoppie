package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

// CheckError catch error and return
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

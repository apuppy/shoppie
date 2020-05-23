package controller

import (
	"encoding/json"
	"shoppie/model"
	"shoppie/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostStaff create staff
func PostStaff(c *gin.Context) {
	body := util.BodyParamsParser(c)
	var staff model.ShopStaff
	err := json.Unmarshal(body, &staff)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}
	err = model.NewShopStaff(staff)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}
	util.JSON(c, gin.H{"data": staff})
}

//GetStaffs get rows from staff
func GetStaffs(c *gin.Context) {
	staffs := model.GetAll(model.SHOPSTAFF)
	util.JSON(c, gin.H{"data": staffs})
}

//GetStaff get staff by id
func GetStaff(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	staff := model.GetByID(model.SHOPSTAFF, id)
	util.JSON(c, gin.H{"data": staff})
}

//DeleteStaff delete staff by id
func DeleteStaff(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	rows, err := model.DeleteByID(model.SHOPSTAFF, id)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}
	util.JSON(c, gin.H{"message": "OK", "rows": rows})
}

//PutStaff update staff
func PutStaff(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	body := util.BodyParamsParser(c)
	var staff model.ShopStaff
	err := json.Unmarshal(body, &staff)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}
	staff.ID = uint32(id)
	rows, err := model.UpdateShopStaff(staff)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}
	util.JSON(c, gin.H{"message": "success", "rows": rows})
}

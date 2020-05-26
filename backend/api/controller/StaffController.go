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

// Password struct for staff generate password
type Password struct {
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// SetPasswd set password for staff
func SetPasswd(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var json Password
	if err := c.ShouldBindJSON(&json); err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
	}
	password := json.Password
	passwordHash, err := util.Hash(password)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
	}
	var staff model.ShopStaff
	rows, err := model.UpdateShopStaffPwd(staff, id, passwordHash)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
	}
	util.JSON(c, gin.H{"message": "success", "rows": rows})
}

// LoginInfo staff login info login_name and password
type LoginInfo struct {
	LoginName string `form:"" json:"login_name" xml:"login_name" binding:"required"`
	Password  string `form:"password" json:"password" xml:"password" binding:"required"`
}

// VerifyPasswd verify password for staff
func VerifyPasswd(c *gin.Context) {
	var loginInfo LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}
	loginName := loginInfo.LoginName
	password := loginInfo.Password

	// get password_hash from database
	passwordHash := model.GetStaffByLoginName(loginName)
	// verify password with password hash stored in DB
	err := util.VerifyPassword(passwordHash, password)
	if err != nil {
		util.JSON(c, gin.H{"message": err.Error()})
		return
	}

	util.JSON(c, gin.H{"message": "success", "verified": "passed"})
}

package controller

import (
	"encoding/json"
	"shoppie/model"
	"shoppie/util"

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
/*
func GetStaff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	staff := model.GetByID(model.SHOPSTAFF, id)
	util.ToJSON(w, staff, http.StatusOK)
}
*/

//DeleteStaff delete staff by id
/*
func DeleteStaff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	rows, err := model.DeleteByID(model.SHOPSTAFF, id)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, rows, http.StatusOK)
}
*/

//PutStaff update staff
/*
func PutStaff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := util.BodyParser(r)
	var staff model.ShopStaff
	err := json.Unmarshal(body, &staff)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	staff.ID = uint32(id)
	rows, err := model.UpdateShopStaff(staff)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, rows, http.StatusOK)
}
*/

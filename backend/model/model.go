package model

import "github.com/jinzhu/gorm"

// table names
const (
	SHOPSTAFF = "shop_staffs"
)

//GetAll get all data from specific table
func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case SHOPSTAFF:
		return db.Order("id ASC").Find(&[]ShopStaff{}).Value
	}
	return nil
}

// GetByID get row by id
func GetByID(table string, id uint64) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case SHOPSTAFF:
		return db.Where("id = ?", id).Find(&ShopStaff{}).Value
	}
	return nil
}

// DeleteByID delete row by id
func DeleteByID(table string, id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	switch table {
	case SHOPSTAFF:
		rs = db.Where("id = ?", id).Delete(&ShopStaff{})
	}
	return rs.RowsAffected, rs.Error
}

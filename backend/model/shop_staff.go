package model

import (
	"fmt"
	"time"
)

// ShopStaff shop_staffs table
type ShopStaff struct {
	ID           uint32     `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Realname     string     `gorm:"type:varchar(15) not null default ''" json:"realname"`
	Nickname     string     `gorm:"type:varchar(15) not null default ''" json:"nickname"`
	PasswordHash string     `gorm:"type:varchar(72) not null default ''" json:"password_hash"`
	Mobile       string     `gorm:"type:varchar(11) not null default ''" json:"mobile"`
	Avatar       string     `gorm:"type:varchar(60) not null default ''" json:"avatar"`
	Motto        string     `gorm:"type:varchar(30) not null default ''" json:"motto"`
	Bio          string     `gorm:"type:varchar(60) not null default ''" json:"bio"`
	Email        string     `gorm:"type:varchar(30) not null default ''" json:"email"`
	CreatedAt    time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"" json:"deleted_at"`
}

// NewShopStaff add new staff
func NewShopStaff(staff ShopStaff) error {
	db := Connect()
	defer db.Close()
	var err error
	if err != nil {
		return nil
	}
	err = db.Create(&staff).Error
	return err
}

// UpdateShopStaff update staff
func UpdateShopStaff(staff ShopStaff) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&staff).Where("id = ?", staff.ID).UpdateColumns(
		map[string]interface{}{
			"nickname": staff.Nickname,
			"email":    staff.Email,
		},
	)
	return rs.RowsAffected, rs.Error
}

// UpdateShopStaffPwd update staff
func UpdateShopStaffPwd(staff ShopStaff, id uint64, passwordHash string) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&staff).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"password_hash": passwordHash,
		},
	)
	return rs.RowsAffected, rs.Error
}

// GetStaffByLoginName get password hash by login name, current is mobile
func GetStaffByLoginName(loginName string) string {
	fmt.Println(loginName)
	db := Connect()
	defer db.Close()
	var hashes []string
	// db.Model(&ShopStaff{}).Pluck("password_hash", &hashes)
	db.Model(&ShopStaff{}).Where("id = ?", 2).Pluck("password_hash", &hashes)
	if len(hashes) > 0 {
		return hashes[0]
	} else {
		return ""
	}
}

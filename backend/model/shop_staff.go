package model

import (
	"time"
)

// ShopStaff shop_staffs table
type ShopStaff struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Realname  string    `gorm:"type:varchar(15) not null default ''" json:"realname"`
	Nickname  string    `gorm:"type:varchar(15) not null default ''" json:"nickname"`
	Mobile    string    `gorm:"type:varchar(11) not null default ''" json:"mobile"`
	Avatar    string    `gorm:"type:varchar(60) not null default ''" json:"avatar"`
	Motto     string    `gorm:"type:varchar(30) not null default ''" json:"motto"`
	Bio       string    `gorm:"type:varchar(60) not null default ''" json:"bio"`
	Email     string    `gorm:"type:varchar(30) not null default ''" json:"email"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

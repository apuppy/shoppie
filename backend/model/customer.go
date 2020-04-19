package model

import (
	"time"
)

// Customer customers table
type Customer struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Nickname  string    `gorm:"type:varchar(20) not null default ''" json:"nickname"`
	Realname  string    `gorm:"type:varchar(20) not null default ''" json:"realname"`
	Mobile    string    `gorm:"type:char(11) not null default ''" json:"mobile"`
	Email     string    `gorm:"type:varchar(30) not null default ''" json:"email"`
	AvatarID  uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"avatar_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

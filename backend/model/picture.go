package model

import (
	"time"
)

// Picture pictures table
type Picture struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	PicType   uint32    `gorm:"type:tinyint(1) unsigned not null default 0" json:"pic_type"`
	Filename  uint32    `gorm:"type:varchar(60) not null default ''" json:"filename"`
	IsDelete  string    `gorm:"type:tinyint(1) unsigned not null default 0" json:"is_delete"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

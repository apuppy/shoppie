package model

import (
	"time"
)

// Bank banks table
type Bank struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	BankName  string    `gorm:"type:varchar(20) not null default ''" json:"bank_name"`
	Logo      string    `gorm:"type:varchar(60) not null default ''" json:"logo"`
	IsDelete  uint32    `gorm:"type:tinyint(1) unsigned not null default 0" json:"is_delete"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

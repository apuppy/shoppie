package model

import (
	"time"
)

// Avatar avatars table
type Avatar struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Filename  string    `gorm:"type:varchar(60) not null default ''" json:"filename"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

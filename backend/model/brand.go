package model

import (
	"time"
)

// Brand brands table
type Brand struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(30) not null default ''" json:"name"`
	Tel       string    `gorm:"type:char(11) not null default ''" json:"tel"`
	Address   string    `gorm:"type:varchar(60) not null default ''" json:"address"`
	Logo      string    `gorm:"type:varchar(60) not null default ''" json:"logo"`
	Slogan    string    `gorm:"type:varchar(30) not null default ''" json:"slogan"`
	IsOpening string    `gorm:"type:tinyint(1) unsigned not null default 0" json:"is_opening"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

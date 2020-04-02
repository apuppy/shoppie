package model

import (
	"time"
)

// Brand brand table
type Brand struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(30);not null" json:"name"`
	Tel       string    `gorm:"type:char(11);not null" json:"tel"`
	Address   string    `gorm:"type:varchar(60);not null" json:"address"`
	Logo      string    `gorm:"type:varchar(60);not null" json:"logo"`
	Slogan    string    `gorm:"type:varchar(30);not null" json:"slogan"`
	IsOpening string    `gorm:"type:tinyint(1);not null" json:"is_opening"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

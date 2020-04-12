package model

import (
	"time"
)

// Shop shop table
type Shop struct {
	ID        uint32    `gorm:"type:int(11) unsigned;primary_key;auto_increment" json:"id"`
	BrandID   uint32    `gorm:"type:int(11) unsigned;not null" json:"brand.id"`
	Name      string    `gorm:"type:varchar(30);not null" json:"name"`
	SN        string    `gorm:"type:varchar(30) not null" json:"sn"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

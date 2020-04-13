package model

import (
	"time"
)

// ProductCategory product_categories table
type ProductCategory struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(20) not null default ''" json:"name"`
	IsOpening string    `gorm:"type:tinyint(1) unsigned not null default 0" json:"is_opening"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

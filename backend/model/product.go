package model

import (
	"time"
)

// Product products table
type Product struct {
	ID           uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Name         string    `gorm:"type:varchar(30) not null default ''" json:"name"`
	CostPrice    uint32    `gorm:"type:int(11) unsigned not null default 0" json:"cost_price"`
	SellingPrice uint32    `gorm:"type:int(11) unsigned not null default 0" json:"selling_price"`
	LeftCount    uint32    `gorm:"type:int(11) unsigned not null default 0" json:"left_count"`
	SoldCount    uint32    `gorm:"type:int(11) unsigned not null default 0" json:"sold_count"`
	IsOpening    uint32    `gorm:"type:tinyint(1) unsigned not null default 0" json:"is_opening"`
	CreatedAt    time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt    time.Time `gorm:"" json:"deleted_at"`
}

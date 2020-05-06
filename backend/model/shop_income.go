package model

import (
	"time"
)

// ShopIncome shop_incomes table
type ShopIncome struct {
	ID              uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	ShopID          uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"shop_id"`
	TotalIncome     uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"total_income"`
	LeftIncome      uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"left_income"`
	WithdrawnIncome uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"withdrawn_income"`
	CreatedAt       time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt       time.Time `gorm:"" json:"deleted_at"`
}

package model

import (
	"time"
)

// ShopIncomeLog shop_income_logs table
type ShopIncomeLog struct {
	ID                 uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	CustomerPayOrderID uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_pay_order_id"`
	Price              uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"price"`
	CreatedAt          time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt          time.Time `gorm:"" json:"deleted_at"`
}

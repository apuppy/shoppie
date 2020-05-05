package model

import (
	"time"
)

// CustomerPayOrderItem customer_pay_order_items table
type CustomerPayOrderItem struct {
	ID                 uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	CustomerID         uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_id"`
	CustomerPayOrderID uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_pay_order_id"`
	CustomerOrderID    uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_order_id"`
	PayPrice           uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"pay_price"`
	CreatedAt          time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt          time.Time `gorm:"" json:"deleted_at"`
}

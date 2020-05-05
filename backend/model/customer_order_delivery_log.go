package model

import (
	"time"
)

// CustomerOrderDeliveryLog customer_order_delivery_logs table
type CustomerOrderDeliveryLog struct {
	ID              uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	CustomerID      uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_id"`
	CustomerOrderID uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_order_id"`
	AddressName     uint32    `gorm:"type:varchar(50) not null default ''" json:"address_name"`
	Lat             float64   `gorm:"type:decimal(10, 8) not null" json:"lat"`
	Lng             float64   `gorm:"type:decimal(11, 8) not null" json:"lng"`
	Message         string    `gorm:"type:text" json:"message"`
	Remark          string    `gorm:"type:text" json:"remark"`
	CreatedAt       time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt       time.Time `gorm:"" json:"deleted_at"`
}

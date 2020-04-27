package model

import (
	"time"
)

// CustomerPayOrder customer_pay_orders table
type CustomerPayOrder struct {
	ID                uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	CustomerID        uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_id"`
	DeliveryAddressID uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"delivery_address_id"`
	PayChannel        uint32    `gorm:"type:tinyint(1) unsigned not null default  '0'" json:"pay_channel"`
	PayPrice          uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"pay_price"`
	PayProgress       uint32    `gorm:"type:tinyint(1) unsigned not null default '0'" json:"pay_progress"`
	PaidAt            time.Time `gorm:"default:current_timestamp()" json:"paid_at"`
	CreatedAt         time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt         time.Time `gorm:"" json:"deleted_at"`
}

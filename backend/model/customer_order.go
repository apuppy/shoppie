package model

import (
	"time"
)

// CustomerOrder customer_orders table
type CustomerOrder struct {
	ID               uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	CustomerID       uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_id"`
	ProductID        uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"product_id"`
	UnitPrice        uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"unit_price"`
	PurchaseCount    uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"purchase_count"`
	DiscountPrice    uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"discount_price"`
	PayProgress      uint32    `gorm:"type:tinyint(1) unsigned not null default '0'" json:"pay_progress"`
	DeliveryProgress uint32    `gorm:"type:tinyint(1) unsigned not null default '0'" json:"delivery_progress"`
	PayPrice         uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"pay_price"`
	IsDelete         uint32    `gorm:"type:tinyint(1) unsigned not null default '0'" json:"is_delete"`
	CreatedAt        time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt        time.Time `gorm:"" json:"deleted_at"`
}

package model

import (
	"time"
)

// CustomerPayOrderLogs customer_pay_order_logs table
type CustomerPayOrderLogs struct {
	ID                 uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	CustomerID         uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_id"`
	CustomerPayOrderID uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"customer_pay_order_id"`
	CallbackMessage    string    `gorm:"type:text" json:"callback_message"`
	CreatedAt          time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt          time.Time `gorm:"" json:"deleted_at"`
}

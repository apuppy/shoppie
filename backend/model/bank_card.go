package model

import (
	"time"
)

// BankCard bank_cards table
type BankCard struct {
	ID            uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	BankID        uint32    `gorm:"type:int(11) unsigned unsigned not null default '0'" json:"bank_id"`
	CardNumber    string    `gorm:"type:varchar(19) not null default ''" json:"card_number"`
	AccountName   string    `gorm:"type:varchar(20) not null default ''" json:"account_name"`
	AccountMobile string    `gorm:"type:char(11) not null default ''" json:"account_mobile"`
	IsDefault     uint32    `gorm:"type:tinyint(1) unsigned not null default 0" json:"is_default"`
	CreatedAt     time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt     time.Time `gorm:"" json:"deleted_at"`
}

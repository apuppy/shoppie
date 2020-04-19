package model

import (
	"time"
)

// DeliveryAddress delivery_addresses table
type DeliveryAddress struct {
	ID            uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	AddressID     uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"address_id"`
	ProvinceID    uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"province_id"`
	CityID        uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"city_id"`
	DistrictID    uint32    `gorm:"type:int(11) unsigned not null default '0'" json:"district_id"`
	DetailAddress string    `gorm:"type:varchar(60) not null default ''" json:"detail_address"`
	IsDefault     uint32    `gorm:"type:tinyint(1) unsigned not null default '0'" json:"is_default"`
	IsDelete      uint32    `gorm:"type:tinyint(1) unsigned not null default '0'" json:"is_delete"`
	CreatedAt     time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt     time.Time `gorm:"" json:"deleted_at"`
}

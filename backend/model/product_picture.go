package model

import (
	"time"
)

// ProductPicture product_pictures table
type ProductPicture struct {
	ID        uint32    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	ProductID uint32    `gorm:"type:int(11) unsigned not null default 0" json:"product_id"`
	PictureID uint32    `gorm:"type:int(11) unsigned not null default 0" json:"picture_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

package model

// AutoMigration database auto migration
func AutoMigration() {
	db := Connect()
	defer db.Close()

	db.Debug().DropTableIfExists(&Brand{}, &Shop{}, &ShopStaff{}, &ProductCategory{})
	db.Debug().AutoMigrate(&Brand{}, &Shop{}, &ShopStaff{}, &ProductCategory{})
}

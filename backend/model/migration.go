package model

// AutoMigration database auto migration
func AutoMigration() {
	db := Connect()
	defer db.Close()

	db.Debug().DropTableIfExists(&Brand{}, &Shop{})
	db.Debug().AutoMigrate(&Brand{}, &Shop{})
}

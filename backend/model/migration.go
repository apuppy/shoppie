package model

// AutoMigration database auto migration
func AutoMigration() {
	db := Connect()
	defer db.Close()

	db.Debug().DropTableIfExists(&Brand{})
	db.Debug().AutoMigrate(&Brand{})
}

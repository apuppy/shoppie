package model

// AutoMigration database auto migration
func AutoMigration() {
	db := Connect()
	defer db.Close()

	db.Debug().DropTableIfExists(&Brand{}, &Shop{}, &ShopStaff{}, &ProductCategory{}, &Product{}, &ProductDetail{},
		&ProductPicture{}, &Picture{}, &Customer{}, &Avatar{}, &DeliveryAddress{}, &ShoppingCart{}, &CustomerOrder{},
		&CustomerPayOrder{}, &CustomerPayOrderLog{}, &CustomerPayOrderItem{}, &CustomerOrderDeliveryLog{}, &ShopIncome{},
		&ShopIncomeLog{}, &Bank{}, &BankCard{})
	db.Debug().AutoMigrate(&Brand{}, &Shop{}, &ShopStaff{}, &ProductCategory{}, &Product{}, &ProductDetail{},
		&ProductPicture{}, &Picture{}, &Customer{}, &Avatar{}, &DeliveryAddress{}, &ShoppingCart{}, &CustomerOrder{},
		&CustomerPayOrder{}, &CustomerPayOrderLog{}, &CustomerPayOrderItem{}, &CustomerOrderDeliveryLog{}, &ShopIncome{},
		&ShopIncomeLog{}, &Bank{}, &BankCard{})
}

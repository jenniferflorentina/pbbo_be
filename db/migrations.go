package db

import (
	modelPayment "tubespbbo/domains/payment/model"
	modelUser "tubespbbo/domains/user/model"
	modelProduct "tubespbbo/domains/product/model"
	modelExpense "tubespbbo/domains/expense/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(modelUser.User))
	_ = Orm.AutoMigrate(new(modelPayment.PaymentMethod))
	_ = Orm.AutoMigrate(new(modelPayment.Payment))
	_ = Orm.AutoMigrate(new(modelProduct.Product))
	_ = Orm.AutoMigrate(new(modelExpense.Expense))
}

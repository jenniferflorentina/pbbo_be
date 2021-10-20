package db

import (
	modelExpense "tubespbbo/domains/expense/model"
	modelPayment "tubespbbo/domains/payment/model"
	modelProduct "tubespbbo/domains/product/model"
	modelTransaction "tubespbbo/domains/transaction/model"
	modelUser "tubespbbo/domains/user/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(modelUser.User))
	_ = Orm.AutoMigrate(new(modelPayment.PaymentMethod))
	_ = Orm.AutoMigrate(new(modelPayment.Payment))
	_ = Orm.AutoMigrate(new(modelProduct.Product))
	_ = Orm.AutoMigrate(new(modelExpense.Expense))
	_ = Orm.AutoMigrate(new(modelTransaction.Transaction))
	_ = Orm.AutoMigrate(new(modelTransaction.TransactionDetail))
}

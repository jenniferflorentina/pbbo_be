package db

import (
	modelPayment "tubespbbo/domains/payment/model"
	modelUser "tubespbbo/domains/user/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(modelUser.User))
	_ = Orm.AutoMigrate(new(modelPayment.PaymentMethod))
	_ = Orm.AutoMigrate(new(modelPayment.Payment))

}

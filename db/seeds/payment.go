package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedPayment() {
	var pm []model.Payment
	var count int64
	db.Orm.Model(&pm).Count(&count)

	if count > 0 {
		return
	}

	paymentMethods := make([]model.Payment, 3)
	paymentMethods[0] = model.Payment{
		TransactionId:     1,
		PaymentMethodId: 1,
		Status:          "Lunas",
		AccountNumber:      "987654321123",
		Amount:          50000,
	}
	paymentMethods[1] = model.Payment{
		TransactionId:     2,
		PaymentMethodId: 2,
		Status:          "Pending",
		AccountNumber:      "085278633421",
		Amount:          40000,
	}
	paymentMethods[2] = model.Payment{
		TransactionId:     3,
		PaymentMethodId: 3,
		Status:          "Belum Bayar",
		AccountNumber:      "27852684",
		Amount:          300000,
	}

	_ = db.Orm.Create(&paymentMethods)
}

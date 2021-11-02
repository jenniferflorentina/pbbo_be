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
		Id:              1,
		TransaksiId:     1,
		PaymentMethodId: 1,
		Status:          "Lunas",
		NoRekening:      "987654321123",
		Amount:          50000,
	}
	paymentMethods[1] = model.Payment{
		Id:              2,
		TransaksiId:     2,
		PaymentMethodId: 2,
		Status:          "Pending",
		NoRekening:      "085278633421",
		Amount:          40000,
	}
	paymentMethods[2] = model.Payment{
		Id:              3,
		TransaksiId:     3,
		PaymentMethodId: 3,
		Status:          "Belum Bayar",
		NoRekening:      "27852684",
		Amount:          300000,
	}

	_ = db.Orm.Create(&paymentMethods)
}

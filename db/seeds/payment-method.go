package seeds

import (
	"tubespbbo/db"
	"tubespbbo/domains/payment/model"
)

func (s *Seed) SeedPM() {
	var pm []model.PaymentMethod
	var count int64
	db.Orm.Model(&pm).Count(&count)

	if count > 0 {
		return
	}

	paymentMethods := make([]model.PaymentMethod, 3)
	paymentMethods[0] = model.PaymentMethod{
		Nama: "Debit Card",
		Kode: "123123123123",
	}
	paymentMethods[1] = model.PaymentMethod{
		Nama: "Gopay",
		Kode: "076735623421",
	}
	paymentMethods[2] = model.PaymentMethod{
		Nama: "CreditCard",
		Kode: "235462134",
	}

	_ = db.Orm.Create(&paymentMethods)
}

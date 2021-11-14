package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
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
		Name: "Debit Card",
		Code: "123123123123",
	}
	paymentMethods[1] = model.PaymentMethod{
		Name: "Gopay",
		Code: "076735623421",
	}
	paymentMethods[2] = model.PaymentMethod{
		Name: "CreditCard",
		Code: "235462134",
	}

	_ = db.Orm.Create(&paymentMethods)
}

package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedTD() {
	var pm []model.TransactionDetail
	var count int64
	db.Orm.Model(&pm).Count(&count)

	if count > 0 {
		return
	}

	TransactionDetails := make([]model.TransactionDetail, 3)
	TransactionDetails[0] = model.TransactionDetail{
		TransaksiId: 1,
		ProductId:   1,
		Jumlah:      2,
	}
	TransactionDetails[1] = model.TransactionDetail{
		TransaksiId: 2,
		ProductId:   2,
		Jumlah:      2,
	}
	TransactionDetails[2] = model.TransactionDetail{
		TransaksiId: 3,
		ProductId:   3,
		Jumlah:      2,
	}

	_ = db.Orm.Create(&TransactionDetails)
}

package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedExpence() {
	var ex []model.Expense
	var count int64
	db.Orm.Model(&ex).Count(&count)

	if count > 0 {
		return
	}

	expenses := make([]model.Expense, 3)
	expenses[0] = model.Expense{
		UserId:		     3,
		Name:			 "Tridia Enjeliani",
		ReleaseDate:     "2021-11-07",
		Quantity:	     15,
		Category:        0,
		Description:     "Beli pupuk organik",
	}
	expenses[1] = model.Expense{
		UserId:		     3,
		Name:			 "Tridia Enjeliani",
		ReleaseDate:     "2021-10-11",
		Quantity:	     20,
		Category:        1,
		Description:     "Beli stok peralatan menanam",
	}
	expenses[2] = model.Expense{
		UserId:		     4,
		Name:			 "Melvin Sebastian",
		ReleaseDate:     "2021-09-22",
		Quantity:	     45,
		Category:        0,
		Description:     "Beli stok bibit bunga baru",
	}

	_ = db.Orm.Create(&expenses)
}

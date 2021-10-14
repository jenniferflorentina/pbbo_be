package seeds

import (
	"tubespbbo/db"
	"tubespbbo/domains/user/model"
)

func (s *Seed) SeedCountry() {
	var country []model.Country
	count, _ := db.Orm.Count(&country)

	if count > 0 {
		return
	}

	countries := make([]model.Country, 3)
	countries[0] = model.Country{
		Name: "Indonesia",
	}
	countries[1] = model.Country{
		Name: "Germany",
	}
	countries[2] = model.Country{
		Name: "England",
	}

	_, _ = db.Orm.Insert(&countries)
}

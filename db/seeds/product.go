package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedProduct() {
	var product []model.Product
	var count int64
	db.Orm.Model(&product).Count(&count)

	if count > 0 {
		return
	}

	products := make([]model.Product, 3)
	products[0] = model.Product{
		Name:        "Bunga Tulip",
		Code:        "1111",
		Description: "Bunga Tulip adalah bunga asli darinegara Belanda",
		Price:       25000,
		Quantity:    45,
	}
	products[1] = model.Product{
		Name:        "Bunga Mawar",
		Code:        "1112",
		Description: "Bunga Mawar banyak diminati ABG",
		Price:       20000,
		Quantity:    85,
	}
	products[2] = model.Product{
		Name:        "Bunga Anggrek",
		Code:        "1113",
		Description: "Bunga Angggrek banyak diminati ibu-ibu",
		Price:       150000,
		Quantity:    50,
	}

	_ = db.Orm.Create(&products)
}

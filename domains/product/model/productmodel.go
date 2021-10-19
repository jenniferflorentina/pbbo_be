package model

import (
	"tubespbbo/base"
)

type Product struct {
	base.Model      `gorm:"extends"`
	Id              int64
	Nama	        string  `gorm:"varchar(100)"`
	Kode		    string  `gorm:"varchar(100)"`
	Description     string  `gorm:"varchar(255)"`
	Price			float32 `gorm:"double"`
	Quantity		int64
	ImageUrl		string  `gorm:"varchar(255)"`
}

func init() {
}
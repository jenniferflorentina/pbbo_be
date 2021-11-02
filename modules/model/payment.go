package model

import (
	"tubespbbo/base"
)

type Payment struct {
	base.Model      `gorm:"extends"`
	Id              int64
	PaymentMethodId int64
	TransaksiId     int64
	PaymentMethod   PaymentMethod
	Status          string  `gorm:"varchar(100)"`
	NoRekening      string  `gorm:"varchar(100)"`
	Amount          float32 `gorm:"double"`
}

type PaymentMethod struct {
	base.Model `gorm:"extends"`
	Id         int64
	Nama       string    `gorm:"varchar(100)"`
	Kode       string    `gorm:"varchar(100)"`
	Payment    []Payment `gorm:"ForeignKey:PaymentMethodId;references:Id"`
}

func init() {
}

package model

import (
	"tubespbbo/base"
)

type Transaction struct {
	base.Model         `gorm:"extends"`
	Id                 int64
	UserId             int64
	TanggalTransaksi   string              `gorm:"varchar(100)"`
	NoResi             string              `gorm:"varchar(100)"`
	Status             string              `gorm:"varchar(100)"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:TransaksiId;references:Id"`
	Payment            Payment             `gorm:"ForeignKey:TransaksiId"`
}

type TransactionDetail struct {
	base.Model  `gorm:"extends"`
	Id          int64
	TransaksiId int64
	ProductId   int64
	Jumlah      int64
	Product     Product
}

func init() {
}

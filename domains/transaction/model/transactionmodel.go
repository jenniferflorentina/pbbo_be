package model

import (
	"tubespbbo/base"
	modelPayment "tubespbbo/domains/payment/model"
	modelProduct "tubespbbo/domains/product/model"
	modelUser "tubespbbo/domains/user/model"
)

type Transaction struct {
	base.Model       `gorm:"extends"`
	Id               int64
	UserId           int64
	PembayaranId     int64
	TanggalTransaksi string                       `gorm:"varchar(100)"`
	NoResi           string                       `gorm:"varchar(100)"`
	Status           string                       `gorm:"varchar(100)"`
	User             []modelUser.User             `gorm:"ForeignKey:Id;references:IdUser"`
	Pembayaran       []modelPayment.PaymentMethod `gorm:"ForeignKey:Id;references:IdPembayaran"`
}

type TransactionDetail struct {
	base.Model  `gorm:"extends"`
	Id          int64
	TransaksiId int64
	ProdukId    int64
	Jumlah      int64
	Transaksi   []Transaction          `gorm:"ForeignKey:Id;references:IdTransaksi"`
	Produk      []modelProduct.Product `gorm:"ForeignKey:Id;references:IdProduct"`
}

func init() {
}

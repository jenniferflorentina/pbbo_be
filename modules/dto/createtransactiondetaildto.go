package dto

type CreateTransactionDetailDTO struct {
	TransaksiId int64 `json:"idTransaksi" validate:"empty=false"`
	ProdukId    int64 `json:"idProduk" validate:"empty=false"`
	Jumlah      int64 `json:"jumlah" validate:"empty=false"`
}

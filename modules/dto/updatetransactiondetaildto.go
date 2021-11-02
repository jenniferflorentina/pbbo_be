package dto

type UpdateTransactionDetailDTO struct {
	TransaksiId int64 `json:"idTransaksi"`
	ProdukId    int64 `json:"idProduk"`
	Jumlah      int64 `json:"jumlah"`
}

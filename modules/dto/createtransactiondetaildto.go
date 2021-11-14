package dto

type CreateTransactionDetailDTO struct {
	TransactionId int64 `json:"transactionId"`
	ProductId    int64 `json:"productId"`
	Quantity      int64 `json:"quantity"`
}

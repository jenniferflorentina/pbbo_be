package dto

type CreateTransactionDTO struct {
	UserId           int64  `json:"userId"`
	TransactionDate	 string `json:"transactionDate" validate:"empty=false"`
	ReceiptNumber           string `json:"receiptNumber" validate:"empty=false"`
	Status           string `json:"status" validate:"empty=false"`
}

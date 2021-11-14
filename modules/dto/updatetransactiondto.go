package dto

type UpdateTransactionDTO struct {
	UserId           int64  `json:"userId"`
	TransactionDate	 string `json:"transactionDate"`
	ReceiptNumber    string `json:"receiptNumber"`
	Status           string `json:"status"`
}

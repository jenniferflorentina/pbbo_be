package dto

type CreatePaymentDTO struct {
	AccountNumber      string  `json:"accountNumber" validate:"empty=false"`
	PaymentMethodId int64   `json:"paymentMethodId"`
	Amount          float32 `json:"amount"`
	TransactionId     int64   `json:"transactionId"`
}

package dto

type UpdatePaymentDTO struct {
	AccountNumber      string  `json:"accountNumber"`
	PaymentMethodId int64   `json:"paymentMethodId"`
	Amount          float32 `json:"amount"`
}

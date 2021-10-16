package dto

type UpdatePaymentDTO struct {
	NoRekening      string  `json:"noRekening"`
	PaymentMethodId int64   `json:"paymentMethodId"`
	Amount          float32 `json:"amount"`
}

package service

import (
	"tubespbbo/domains/payment/model"
	"tubespbbo/domains/payment/repository"
)

func FindPaymentMethod() (*[]model.PaymentMethod, error) {
	return repository.FindPaymentMethod()
}

func FindOnePaymentMethod(id int64) (*model.PaymentMethod, error) {
	return repository.FindOnePaymentMethod(id)
}

func CreatePaymentMethod(pm *model.PaymentMethod) error {
	return repository.CreatePaymentMethod(pm)
}

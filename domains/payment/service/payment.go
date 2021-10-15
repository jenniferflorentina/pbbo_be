package service

import (
	"tubespbbo/domains/payment/model"
	"tubespbbo/domains/payment/repository"
)

func FindPayment() (*[]model.Payment, error) {
	return repository.FindPayment()
}

func FindOnePayment(id int64) (*model.Payment, error) {
	return repository.FindOnePayment(id)
}

func CreatePayment(payment *model.Payment) error {
	return repository.CreatePayment(payment)
}

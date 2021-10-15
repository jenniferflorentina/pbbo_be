package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/domains/payment/model"
)

func FindPayment() (*[]model.Payment, error) {
	var payments []model.Payment
	result := db.Orm.Find(&payments).Preload("PaymentMethod").Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment methods found")
	}

	return &payments, nil
}

func FindOnePayment(id int64) (*model.Payment, error) {
	var payment model.Payment
	result := db.Orm.First(&payment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment found")
	}

	return &payment, nil
}

func CreatePayment(pm *model.Payment) error {
	result := db.Orm.Create(pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

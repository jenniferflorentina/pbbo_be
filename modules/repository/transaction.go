package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindTransaction() (*[]model.Transaction, error) {
	var transactions []model.Transaction
	result := db.Orm.Find(&transactions).Preload("Payment").Preload("Payment.PaymentMethod").Preload("TransactionDetails").Preload("TransactionDetails.Product").Find(&transactions).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no transaction found")
	}

	return &transactions, nil
}

func FindOneTransaction(id int64) (*model.Transaction, error) {
	var transaction model.Transaction
	result := db.Orm.First(&transaction, id).Preload("Payment").Preload("Payment.PaymentMethod").Preload("TransactionDetails").Preload("TransactionDetails.Product").First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no transaction found")
	}

	return &transaction, nil
}

func CreateTransaction(pm *model.Transaction) error {
	result := db.Orm.Create(pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateTransaction(pm *model.Transaction) (*model.Transaction, error) {
	result := db.Orm.Save(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	pm, _ = FindOneTransaction(pm.Id)
	return pm, nil
}

func DeleteTransaction(pm *model.Transaction) (*model.Transaction, error) {
	result := db.Orm.Delete(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	return pm, nil
}

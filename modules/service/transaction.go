package service

import (
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindTransaction() (*[]model.Transaction, error) {
	return repository.FindTransaction()
}

func FindOneTransaction(id int64) (*model.Transaction, error) {
	return repository.FindOneTransaction(id)
}

func CreateTransaction(transaction *model.Transaction) error {
	return repository.CreateTransaction(transaction)
}

func UpdateTransaction(updateDto *dto.UpdateTransactionDTO, id int64) (*model.Transaction, error) {
	pm, err := repository.FindOneTransaction(id)
	if err != nil {
		return nil, err
	}
	if updateDto.UserId != 0 {
		pm.UserId = updateDto.UserId
	}
	if updateDto.TanggalTransaksi != "" {
		pm.TanggalTransaksi = updateDto.TanggalTransaksi
	}
	if updateDto.NoResi != "" {
		pm.NoResi = updateDto.NoResi
	}
	if updateDto.Status != "" {
		pm.Status = updateDto.Status
	}

	pm, err = repository.UpdateTransaction(pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeleteTransaction(id int64) (*model.Transaction, error) {
	transaction, err := repository.FindOneTransaction(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteTransaction(transaction)
}

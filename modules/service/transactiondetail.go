package service

import (
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindTransactionDetail() (*[]model.TransactionDetail, error) {
	return repository.FindTransactionDetail()
}

func FindOneTransactionDetail(id int64) (*model.TransactionDetail, error) {
	return repository.FindOneTransactionDetail(id)
}

func CreateTransactionDetail(pm *model.TransactionDetail) error {
	return repository.CreateTransactionDetail(pm)
}

func UpdateTransactionDetail(updateDto *dto.UpdateTransactionDetailDTO, id int64) (*model.TransactionDetail, error) {
	pm, err := repository.FindOneTransactionDetail(id)
	if err != nil {
		return nil, err
	}
	if updateDto.TransaksiId != 0 {
		pm.TransaksiId = updateDto.TransaksiId
	}
	if updateDto.ProdukId != 0 {
		pm.ProductId = updateDto.ProdukId
	}
	if updateDto.Jumlah != 0 {
		pm.Jumlah = updateDto.Jumlah
	}
	pm, err = repository.UpdateTransactionDetail(pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeleteTransactionDetail(id int64) (*model.TransactionDetail, error) {
	pm, err := repository.FindOneTransactionDetail(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteTransactionDetail(pm)
}

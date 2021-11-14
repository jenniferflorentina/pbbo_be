package service

import (
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
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

func UpdatePaymentMethod(updateDto *dto.UpdatePaymentMethodDTO, id int64) (*model.PaymentMethod, error) {
	pm, err := repository.FindOnePaymentMethod(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Code != "" {
		pm.Code = updateDto.Code
	}
	if updateDto.Name != "" {
		pm.Name = updateDto.Name
	}
	pm, err = repository.UpdatePaymentMethod(pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeletePaymentMethod(id int64) (*model.PaymentMethod, error) {
	pm, err := repository.FindOnePaymentMethod(id)
	if err != nil {
		return nil, err
	}
	return repository.DeletePaymentMethod(pm)
}

package service

import (
	"tubespbbo/domains/user/dto"
	"tubespbbo/domains/user/model"
	"tubespbbo/domains/user/repository"
)

func FindUser() (*[]model.User, error) {
	return repository.FindUser()
}

func FindOneUser(id int64) (*model.User, error) {
	return repository.FindOneUser(id)
}

func CreateUser(user *model.User) error {
	return repository.CreateUser(user)
}


func UpdateUser(updateDto *dto.UpdateUserDTO, id int64) (*model.User, error) {
	user, err := repository.FindOneUser(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Alamat != "" {
		user.Alamat = updateDto.Alamat
	}
	if updateDto.Telepon != "" {
		user.Telepon = updateDto.Telepon
	}
	if updateDto.Username != "" {
		user.Username = updateDto.Username
	}
	user, err = repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int64) (*model.User, error) {
	user, err := repository.FindOneUser(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteUser(user)
}
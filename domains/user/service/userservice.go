package service

import (
	"tubespbbo/domains/user/model"
	"tubespbbo/domains/user/repository"
)

func Find() (*[]model.User, error) {
	return repository.Find()
}

func Create(user *model.User) error {
	return repository.Create(user)
}

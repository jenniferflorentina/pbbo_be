package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/domains/user/model"
)

func Find() (*[]model.User, error) {
	var users []model.User
	err := db.Orm.Find(&users)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("no users found")
	}

	return &users, nil
}

func Create(user *model.User) error {
	_, err := db.Orm.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

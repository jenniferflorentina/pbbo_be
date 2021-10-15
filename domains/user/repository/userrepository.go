package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/domains/user/model"
)

func Find() (*[]model.User, error) {
	var users []model.User
	result := db.Orm.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no users found")
	}

	return &users, nil
}

func Create(user *model.User) error {
	result := db.Orm.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

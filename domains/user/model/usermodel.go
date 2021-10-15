package model

import (
	"tubespbbo/base"
)

type User struct {
	base.Model `gorm:"extends"`
	Id         int64
	Email      string `gorm:"varchar(100) unique=true"`
	Password   string `gorm:"varchar(100)"`
	FirstName  string `gorm:"varchar(30)"`
	LastName   string `gorm:"varchar(30)"`
	Role       string `gorm:"varchar(30)"`
}

func init() {
}

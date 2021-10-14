package model

import (
	"tubespbbo/base"
)

type User struct {
	base.Model `xorm:"extends"`
	CountryId  int64
	Email      string `xorm:"varchar(100) unique=true"`
	Password   string `xorm:"varchar(100)"`
	FirstName  string `xorm:"varchar(30)"`
	LastName   string `xorm:"varchar(30)"`
	Role       string `xorm:"varchar(30)"`
}

func init() {
}

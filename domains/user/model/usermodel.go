package model

import (
	"tubespbbo/base"
)

type User struct {
	base.Model `gorm:"extends"`
	Id         int64
	Username   string `gorm:"varchar(100)"`
	Password   string `gorm:"varchar(255)"`
	Nama       string `gorm:"varchar(40)"`
	Dob        string `gorm:"varchar(10)"`
	Tipe       int
	Alamat     string `gorm:"varchar(50)"`
	Telepon    string `gorm:"varchar(15)"`
	Email      string `gorm:"varchar(100) unique=true"`
}

func init() {
}

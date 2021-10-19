package model

import (
	"tubespbbo/base"
)

type Expense struct {
	base.Model      `gorm:"extends"`
	Id              int64
	UserId 			int64
	Nama	        string  `gorm:"varchar(100)"`
	ReleaseDate		string  `gorm:"varchar(10)"`
	Quantity        float32 `gorm:"double"`
	Category		int
	Description	    string  `gorm:"varchar(255)"`
}

func init() {
}
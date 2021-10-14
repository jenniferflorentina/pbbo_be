package model

import (
	"tubespbbo/base"
)

type Company struct {
	base.Model `xorm:"extends"`
	Name       string `xorm:"varchar(30)"`
}

func init() {
}

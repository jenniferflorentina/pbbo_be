package base

import "time"

type Model struct {
	Id        int64
	CreatedAt time.Time `xorm:"created"`
	CreatedBy int64
	UpdatedAt time.Time `xorm:"updated"`
	UpdatedBy int64
	DeletedAt time.Time `xorm:"deleted"`
	DeletedBy int64
}

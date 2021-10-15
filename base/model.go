package base

import "time"

type Model struct {
	CreatedAt time.Time `gorm:"created"`
	CreatedBy int64
	UpdatedAt time.Time `gorm:"updated"`
	UpdatedBy int64
	DeletedAt time.Time `gorm:"deleted"`
	DeletedBy int64
}

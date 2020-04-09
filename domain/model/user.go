package model

import "time"

type User struct {
	ID         uint64 `gorm:"primary_key"`
	Name       string
	Updatetime time.Time
	Createtime time.Time
}

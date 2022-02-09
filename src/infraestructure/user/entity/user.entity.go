package entity

import (
	"time"

	"gorm.io/gorm"
)

// type Password string

// func (p Password) Scan() error {
// 	return "jfhhhgherhg"
// }

type User struct {
	gorm.Model
	Id            int `gorm:"primaryKey"`
	Name          string
	Password      string
	Creation_date time.Time
}

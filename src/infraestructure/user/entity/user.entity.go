package entity

import (
	"time"

	"github.com/bigpigeon/toyorm"
)

// type Password string

// func (p Password) Scan() error {
// 	return "jfhhhgherhg"
// }

type User struct {
	toyorm.ModelDefault
	id            int `toyorm:"primary key;auto_increment"`
	name          string
	password      string
	creation_date time.Time
}

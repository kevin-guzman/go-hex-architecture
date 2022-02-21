package model

import (
	"fmt"
	"time"
)

const (
	MIN_PASSWORD_LENGTH int    = 6
	ERR_LENGTH          string = "The leng of the password is incorrect"
)

type User struct {
	Name          string
	Password      string
	Creation_date time.Time
}

func NewUser(name, password string, creation_date time.Time) (*User, error) {
	if len(password) < MIN_PASSWORD_LENGTH {
		return nil, fmt.Errorf(ERR_LENGTH)
	}
	return &User{
		Name:          name,
		Password:      password,
		Creation_date: creation_date,
	}, nil
}

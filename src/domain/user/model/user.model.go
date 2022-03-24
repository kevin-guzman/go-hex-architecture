package model

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
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

func NewUser(name, password string, creation_date string) (*User, *errors.ErrorCore) {
	if len(password) < MIN_PASSWORD_LENGTH {
		err := fmt.Errorf(ERR_LENGTH)
		return nil, errors.NewErrorCore(err, "/src/domain/user/model/user.model.go", err.Error(), 500)
	}
	date, err := time.Parse("2006-01-02T15:04:05Z", creation_date)
	if err != nil {
		err := fmt.Errorf("Error formating the date %s", creation_date)
		return nil, errors.NewErrorCore(err, "/src/domain/user/model/user.model.go", err.Error(), 400)
	}
	return &User{
		Name:          name,
		Password:      password,
		Creation_date: date,
	}, nil
}

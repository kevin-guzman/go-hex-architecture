package command

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
	"net/http"
	"time"
)

type HandlerRegisterUser struct {
	serviceRegisterUser service.ServiceRegisterUser
}

func NewHandlerRegisterUser(sru *service.ServiceRegisterUser) *HandlerRegisterUser {
	return &HandlerRegisterUser{
		serviceRegisterUser: *sru,
	}
}

func (handleRegisterUser *HandlerRegisterUser) Run(commandRU CommandRegisterUser) interface{} {
	date, err := time.Parse("2006-01-02T15:04:05Z", commandRU.CreationDate)
	if err != nil {
		err := fmt.Errorf("Error formating the date %s", commandRU.CreationDate)
		return errors.NewErrorCore(err, "", err.Error(), http.StatusBadRequest)
	}
	user, err := model.NewUser(commandRU.Name, commandRU.Password, date)
	if err != nil {
		return errors.NewErrorCore(err, "", err.Error(), http.StatusBadRequest)
	}
	return handleRegisterUser.serviceRegisterUser.Run(*user)
}

package command

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
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

func (hru *HandlerRegisterUser) Run(commandRU CommandRegisterUser) (string, error, int) {
	date, err := time.Parse("2006-01-02T15:04:05Z", commandRU.CreationDate)
	if err != nil {
		return "nil", fmt.Errorf("Error formating the date %s", commandRU.CreationDate), 400
	}
	user, err := model.NewUser(commandRU.Name, commandRU.Password, date)
	if err != nil {
		return "", err, 500
	}
	message, err, status := hru.serviceRegisterUser.Run(*user)
	return message, err, status
}

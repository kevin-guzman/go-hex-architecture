package command

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
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
	user, err := model.NewUser(commandRU.Name, commandRU.Password, commandRU.CreationDate)
	if err != nil {
		return "", err, 500
	}
	message, err, status := hru.serviceRegisterUser.Run(*user)
	return message, err, status
}

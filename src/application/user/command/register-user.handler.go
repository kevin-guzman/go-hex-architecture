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

func (hru *HandlerRegisterUser) Run(commandRU CommandRegisterUser) (string, error) {
	user, err := model.NewUser(commandRU.Name, commandRU.Password, commandRU.CreationDate)
	if err != nil {
		return "", err
	}
	message, err := hru.serviceRegisterUser.Run(*user)
	return message, err
}

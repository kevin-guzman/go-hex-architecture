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

func (handleRegisterUser *HandlerRegisterUser) Run(commandRU CommandRegisterUser) interface{} {
	user, err := model.NewUser(commandRU.Name, commandRU.Password, commandRU.CreationDate)
	if err != nil {
		return err
	}
	return handleRegisterUser.serviceRegisterUser.Run(*user)
}

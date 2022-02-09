package controller

import (
	"golang-gingonic-hex-architecture/src/application/user/command"
	"golang-gingonic-hex-architecture/src/application/user/query"
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
)

type ControllerUser struct {
	handlerRegisterUser command.HandlerRegisterUser
	handlerListUsers    query.HandlerListUsers
}

func NewControllerUser(hru command.HandlerRegisterUser, hlu query.HandlerListUsers) *ControllerUser {
	return &ControllerUser{
		handlerRegisterUser: hru,
		handlerListUsers:    hlu,
	}
}

func (cu *ControllerUser) Create(command command.CommandRegisterUser) (string, error) {
	return cu.handlerRegisterUser.Run(command)
}

func (cu *ControllerUser) List() []*dto.UserDto {
	return cu.handlerListUsers.Run()
}

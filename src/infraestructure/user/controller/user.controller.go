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

func NewControllerUser(handleRegisterUser command.HandlerRegisterUser, handleListUser query.HandlerListUsers) *ControllerUser {
	return &ControllerUser{
		handlerRegisterUser: handleRegisterUser,
		handlerListUsers:    handleListUser,
	}
}

func (cu *ControllerUser) Create(command command.CommandRegisterUser) interface{} {
	return cu.handlerRegisterUser.Run(command)
}

func (cu *ControllerUser) List() []*dto.UserDto {
	return cu.handlerListUsers.Run()
}

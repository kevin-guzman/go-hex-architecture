package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
)

var (
	errTrace       string = "/src/domain/user/service/register-user.service.go"
	successMessage string = "User has succesfully created!"
)

type ServiceRegisterUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceRegisterUser(UserR repository.RepositoryUser) *ServiceRegisterUser {
	return &ServiceRegisterUser{
		userRepository: UserR,
	}
}

func (sru *ServiceRegisterUser) Run(user model.User) interface{} {
	existUserName, err := sru.userRepository.ExistUserName(user.Name)
	if err != nil {
		return errors.NewErrorPort(err)
	}
	if existUserName {
		err := fmt.Errorf("The username %s already exist", user.Name)
		return errors.NewErrorUserAlreadyExist(err, err.Error())
	}

	err = sru.userRepository.Save(user)
	if err != nil {
		fmt.Println(errTrace)
		return errors.NewErrorPort(err)
	}

	return successMessage
}

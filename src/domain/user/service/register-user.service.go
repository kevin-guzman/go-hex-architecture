package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
)

type ServiceRegisterUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceRegisterUser(UserR repository.RepositoryUser) *ServiceRegisterUser {
	return &ServiceRegisterUser{
		userRepository: UserR,
	}
}

func (sru *ServiceRegisterUser) Run(user model.User) (string, error) {
	existUserName, err := sru.userRepository.ExistUserName(user.Name)
	if err != nil {
		return "", err
	}
	if existUserName {
		return "", fmt.Errorf("The username %s already exist", user.Name)
	}

	err = sru.userRepository.Save(user)
	if err != nil {
		return "", err
	}

	return "User has succesfully created!", nil
}

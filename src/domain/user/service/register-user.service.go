package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"net/http"
)

var (
	errTrace       string = "This error has ocurred in register-user.service.go"
	internalError  string = "Internal server error"
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

func (sru *ServiceRegisterUser) Run(user model.User) (string, error, int) {
	existUserName, err := sru.userRepository.ExistUserName(user.Name)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}
	if existUserName {
		err := fmt.Errorf("The username %s already exist", user.Name)
		return "", errors.NewErrorCore(err, errTrace, err.Error()).PublicError(), http.StatusInternalServerError
	}

	err = sru.userRepository.Save(user)
	if err != nil {
		fmt.Println(errTrace)
		return "", errors.NewErrorCore(err, errTrace, internalError).PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}

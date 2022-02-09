package service

import (
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

func GetServiceRegisterUser(ru repository.RepositoryUser) *service.ServiceRegisterUser {
	return service.NewServiceRegisterUser(ru)
}

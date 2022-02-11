package repository

import (
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/user/port/repository"
	classRepository "golang-gingonic-hex-architecture/src/infraestructure/user/adaptor/repository"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceRepository.RepositoryUser

var GetRepositoryUser = func(conn *gorm.DB) *interfaceRepository.RepositoryUser {
	once.Do(func() {
		ru := classRepository.NewRepositoryUserPostgreSql(conn)
		iru := interfaceRepository.RepositoryUser(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}

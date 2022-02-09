package repository

import (
	// "golang-gingonic-hex-architecture/src/application/user/command"
	// "golang-gingonic-hex-architecture/src/application/user/query"
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/user/port/repository"
	classRepository "golang-gingonic-hex-architecture/src/infraestructure/user/adaptor/repository"
	"sync"

	toy "github.com/bigpigeon/toyorm"
)

var once sync.Once
var instance *interfaceRepository.RepositoryUser

func GetRepositoryUser(conn toy.Toy) *interfaceRepository.RepositoryUser {
	once.Do(func() {
		ru := classRepository.NewRepositoryUserPostgreSql(conn)
		iru := interfaceRepository.RepositoryUser(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance

	// var in *interfaceRepository.RepositoryUser
	// in = new(interfaceRepository.RepositoryUser)
	// *in = classRepository.NewRepositoryUserPostgreSql(conn)

	// r:=classRepository.RepositoryUserPostgreSql{}
	// u:=classRepository.NewRepositoryUserPostgreSql(conn)
	// *interfaceRepository.RepositoryUser = u
	// var _ interfaceRepository.RepositoryUser = (u)()
	// var _ interfaceRepository.RepositoryUser = (classRepository.NewRepositoryUserPostgreSql(conn))(nil)
	// var _ interfaceRepository.RepositoryUser = (*classRepository.RepositoryUserPostgreSql)(nil)

}

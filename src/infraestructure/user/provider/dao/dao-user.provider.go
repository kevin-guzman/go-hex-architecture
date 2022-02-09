package dao

import (
	// "golang-gingonic-hex-architecture/src/application/user/command"
	// "golang-gingonic-hex-architecture/src/application/user/query"
	interfaceDao "golang-gingonic-hex-architecture/src/domain/user/port/dao"
	classDao "golang-gingonic-hex-architecture/src/infraestructure/user/adaptor/dao"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceDao.DaoUser

func GetDaoUser(conn *gorm.DB) *interfaceDao.DaoUser {
	once.Do(func() {
		ru := classDao.NewDaoUserPostgreSql(conn)
		iru := interfaceDao.DaoUser(ru)
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

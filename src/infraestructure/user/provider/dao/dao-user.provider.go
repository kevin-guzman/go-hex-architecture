package dao

import (
	interfaceDao "golang-gingonic-hex-architecture/src/domain/user/port/dao"
	classDao "golang-gingonic-hex-architecture/src/infraestructure/user/adaptor/dao"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceDao.DaoUser

var GetDaoUser = func(conn *gorm.DB) *interfaceDao.DaoUser {
	once.Do(func() {
		ru := classDao.NewDaoUserPostgreSql(conn)
		iru := interfaceDao.DaoUser(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}

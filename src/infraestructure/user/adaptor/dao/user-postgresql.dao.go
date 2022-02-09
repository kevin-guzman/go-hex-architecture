package dao

import (
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"

	toy "github.com/bigpigeon/toyorm"
)

type DaoUserPostgreSql struct {
	daoUser *toy.ToyBrick
}

func NewDaoUserPostgreSql(conn toy.Toy) *DaoUserPostgreSql {
	return &DaoUserPostgreSql{
		daoUser: conn.Model(&entity.User{}),
	}
}

func (dup *DaoUserPostgreSql) List() []*dto.UserDto {
	var users []*dto.UserDto
	_, _ = dup.daoUser.Template("SELECT u.name FROM USER u").Find(&users)
	return users
}

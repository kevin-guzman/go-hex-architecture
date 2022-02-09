package query

import (
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/domain/user/port/dao"
)

type HandlerListUsers struct {
	daoUser dao.DaoUser
}

func NewHandlerListUsers(daoU dao.DaoUser) *HandlerListUsers {
	return &HandlerListUsers{
		daoUser: daoU,
	}
}

func (hlu *HandlerListUsers) Run() []*dto.UserDto {
	return hlu.daoUser.List()
}

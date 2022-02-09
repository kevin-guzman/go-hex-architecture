package dao

import "golang-gingonic-hex-architecture/src/application/user/query/dto"

type DaoUser interface {
	List() []*dto.UserDto
}

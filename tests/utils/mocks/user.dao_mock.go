package mocks

import (
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/user/adaptor/dao"

	"github.com/stretchr/testify/mock"
)

type MockDaoUser struct {
	dao.DaoUserPostgreSql
	mock.Mock
}

func (m *MockDaoUser) List() []*dto.UserDto {
	_mc_ret := m.Called()

	var _r0 []*dto.UserDto

	if _rfn, ok := _mc_ret.Get(0).(func() []*dto.UserDto); ok {
		_r0 = _rfn()
	} else {
		if _mc_ret.Get(0) != nil {
			_r0 = _mc_ret.Get(0).([]*dto.UserDto)
		}
	}

	return _r0
}

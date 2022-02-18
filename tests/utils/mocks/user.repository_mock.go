package mocks

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/infraestructure/user/adaptor/repository"

	"github.com/stretchr/testify/mock"
)

type MockRepositoryUser struct {
	repository.RepositoryUserPostgreSql
	mock.Mock
}

func (m *MockRepositoryUser) ExistUserName(name string) (bool, error) {
	args := m.Called(name)
	var returned bool
	if len(args) > 1 {
		returned = args.Get(0).(bool)
	}
	return returned, nil
}

func (m *MockRepositoryUser) Save(user model.User) error {
	// _mc_ret := m.Called()

	// var _r0 error

	// if _rfn, ok := _mc_ret.Get(0).(func() error); ok {
	// 	_r0 = _rfn()
	// } else {
	// 	if _mc_ret.Get(0) != nil {
	// 		_r0 = _mc_ret.Get(0).(error)
	// 	}
	// }
	// return _r0

	args := m.Called(user)
	var returned error
	if len(args) > 1 {
		returned = args.Get(0).(error)
	}
	return returned
}

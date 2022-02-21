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
	args := m.Called(user)
	var returned error
	if len(args) > 1 {
		returned = args.Get(0).(error)
	}
	return returned
}

package service_test

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
	"golang-gingonic-hex-architecture/tests/utils/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceSuccessCretionOfUser(t *testing.T) {
	assert := require.New(t)

	repositoryUser := &mocks.MockRepositoryUser{}
	repositoryUser.On("ExistUserName").Return(false, nil)
	repositoryUser.On("Save").Return(nil)

	usr := model.User{Name: "Juan"}
	serviceRegisterUserStub := service.NewServiceRegisterUser(repositoryUser)
	msg, err, code := serviceRegisterUserStub.Run(usr)

	repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
	repositoryUser.AssertCalled(t, "ExistUserName")
	repositoryUser.AssertNumberOfCalls(t, "Save", 1)
	repositoryUser.AssertCalled(t, "Save")
	repositoryUser.AssertExpectations(t)

	assert.True(msg == "User has succesfully created!")
	assert.True(err == nil)
	assert.True(code == 200)
}

func TestServiceIfAlreadyExistUser(t *testing.T) {
	assert := require.New(t)

	repositoryUser := &mocks.MockRepositoryUser{}
	repositoryUser.On("ExistUserName").Return(true, nil)
	repositoryUser.On("Save").Return(nil)

	usr := model.User{Name: "Juan"}
	serviceRegisterUserStub := service.NewServiceRegisterUser(repositoryUser)
	msg, err, code := serviceRegisterUserStub.Run(usr)

	repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
	repositoryUser.AssertCalled(t, "ExistUserName")
	repositoryUser.AssertNumberOfCalls(t, "Save", 0)

	assert.True(msg == "")
	assert.True(err.Error() == "The username "+usr.Name+" already exist")
	assert.True(code == 500)
}

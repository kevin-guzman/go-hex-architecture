package service_test

import (
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
	"golang-gingonic-hex-architecture/tests/utils/mocks"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	repositoryUser mocks.MockRepositoryUser
	usr            = model.User{Name: "Juan"}
	t              *testing.T
)

var _ = Describe("Service create user", func() {
	BeforeSuite(func() {
		t = tReference
		errors.NewErrorCore = func(err error, trace, message string, status int) *errors.ErrorCore {
			return &errors.ErrorCore{
				Err:     err,
				Trace:   trace,
				Message: message,
				Status:  status,
			}
		}
	})

	BeforeEach(func() {
		repositoryUser = mocks.MockRepositoryUser{}
	})

	It("Should create a user", func() {
		repositoryUser.On("ExistUserName", usr.Name).Return(false, nil)
		repositoryUser.On("Save", usr).Return(nil)

		serviceRegisterUserStub := service.NewServiceRegisterUser(&repositoryUser)
		response := serviceRegisterUserStub.Run(usr)

		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
		repositoryUser.AssertCalled(t, "ExistUserName", usr.Name)
		repositoryUser.AssertNumberOfCalls(t, "Save", 1)
		repositoryUser.AssertCalled(t, "Save", usr)
		repositoryUser.AssertExpectations(t)

		Expect(response).To(Equal("User has succesfully created!"))
	})

	It("If user already exists", func() {
		repositoryUser.On("ExistUserName", usr.Name).Return(true, nil)

		serviceRegisterUserStub := service.NewServiceRegisterUser(&repositoryUser)
		response := serviceRegisterUserStub.Run(usr)

		var err *errors.ErrorCore = response.(*errors.ErrorCore)

		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
		repositoryUser.AssertCalled(t, "ExistUserName", usr.Name)
		repositoryUser.AssertNumberOfCalls(t, "Save", 0)
		repositoryUser.AssertExpectations(t)

		Expect(err.PublicError().Error()).To(Equal("The username " + usr.Name + " already exist"))
		Expect(err.Status).To(BeIdenticalTo(http.StatusInternalServerError))
	})
})

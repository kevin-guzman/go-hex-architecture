package model_test

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	It("Should fail with a password less than 6", func() {
		_, err := model.NewUser("JUan", "1234", time.Now())
		Expect(err).Error()
		Expect(err.Error()).To(Equal("The leng of the password is incorrect"))
	})

	It("Should create a user", func() {
		expectUser := model.User{Name: "JUan", Password: "123ss3r4", Creation_date: time.Now()}
		usr, err := model.NewUser(expectUser.Name, expectUser.Password, expectUser.Creation_date)
		Expect(err).To(BeNil())
		Expect(&expectUser).To(Equal(usr))
	})
})

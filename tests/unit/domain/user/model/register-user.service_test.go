package model_test

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	It("Should fail with a password less than 6", func() {
		_, err := model.NewUser("JUan", "1234", time.Now().String())

		Expect(err).Error()
		Expect(err.PublicError().Error()).To(Equal("The leng of the password is incorrect"))
		Expect(err.Status).To(BeIdenticalTo(http.StatusBadRequest))
	})

	It("Should fail with a invalid date", func() {
		const invalidDate string = "wdwewfd"
		_, err := model.NewUser("JUan", "dwde1234", invalidDate)

		Expect(err).Error()
		Expect(err.PublicError().Error()).To(Equal(fmt.Sprintf("Error formating the date %s", invalidDate)))
		Expect(err.Status).To(BeIdenticalTo(http.StatusBadRequest))
	})

	It("Should create a user", func() {
		input := "2017-08-31T15:04:05Z"
		layout := "2006-01-02T15:04:05Z"
		date, _ := time.Parse(layout, input)
		expectUser := model.User{Name: "JUan", Password: "123ss3r4", Creation_date: date}
		usr, err := model.NewUser(expectUser.Name, expectUser.Password, expectUser.Creation_date.Format(layout))

		Expect(err).To(BeNil())
		Expect(&expectUser).To(Equal(usr))
	})
})

package models_test

import (
	. "choreboard/models"
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Model", func() {
	var (
		user       User
		err        error
		jsonString string
	)

	Describe("loading from JSON", func() {
		BeforeEach(func() {
			jsonString = `{
            "id": "1234",
            "first_name": "John",
            "last_name": "Smith",
			"email": "test@test.com",
			"password": "123",
			"created": "2014-06-29T20:56:15Z"
        }`

		})

		JustBeforeEach(func() {
			user = User{}
			err = json.Unmarshal([]byte(jsonString), &user)
		})

		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(user.Id).To(Equal("1234"))
				Expect(user.FirstName).To(Equal("John"))
				Expect(user.LastName).To(Equal("Smith"))
				Expect(user.Email).To(Equal("test@test.com"))
				Expect(user.Password).To(Equal("123"))
				Expect(*user.Created).To(Equal(time.Date(2014, time.June, 29, 20, 56, 15, 0, time.UTC)))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				jsonString = `{
		            "id": 1234,
		            "first_name": "John",
		            "last_name": "Smith",
					"email": "test@test.com",
					"password": "123",
					"created": "2014-06-29T20:56:15Z"
		        }`
			})

			It("should return the zero-value for the user", func() {
				Expect(user).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Describe("encyrpting and decrypting password", func() {
		BeforeEach(func() {
			now := time.Now()
			user = User{
				"id123",
				"Johhny",
				"Five",
				"testing@email.com",
				"",
				&now,
			}
			err = user.SetEncryptPassword("supersecretpassword")
		})

		Context("when password is encrypted succesfully", func() {
			It("should not equal plain text password", func() {
				Expect(user.Password).ShouldNot(Equal("supersecretpassword"))
			})
		})

		Context("when the same password is encrypted more than once", func() {
			It("should not be the same encyrpted value", func() {
				firstEncryption := user.Password
				err = user.SetEncryptPassword("supersecretpassword")
				secondEncryption := user.Password
				Expect(firstEncryption).ShouldNot(Equal(secondEncryption))
			})
		})

		Context("when password is decrypted succesfully", func() {
			It("should equal plain text password", func() {
				pwd, _ := user.GetDecryptPassword()
				Expect(pwd).Should(Equal("supersecretpassword"))
			})
		})

	})
})

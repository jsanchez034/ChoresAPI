package models_test

import (
	. "choreboard/models"
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chore Model", func() {
	var (
		chore      Chore
		err        error
		jsonString string
	)

	BeforeEach(func() {
		jsonString = `{
            "Id": "123",
            "Name": "Do dishes",
            "Description": "Please do the dishes",
			"Created": "2014-06-29T20:56:15Z"
        }`

	})

	JustBeforeEach(func() {
		chore = Chore{}
		err = json.Unmarshal([]byte(jsonString), &chore)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(chore.Id).To(Equal("123"))
				Expect(chore.Name).To(Equal("Do dishes"))
				Expect(chore.Description).To(Equal("Please do the dishes"))
				Expect(*chore.Created).To(Equal(time.Date(2014, time.June, 29, 20, 56, 15, 0, time.UTC)))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				jsonString = `{
		            "Id": 123,
		            "Name": "Do dishes",
		            "Description": "Please do the dishes",
					"Created": "2014-06-29T20:56:15Z"
                }`
			})

			It("should return the zero-value for the chore", func() {
				Expect(chore).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

	})
})

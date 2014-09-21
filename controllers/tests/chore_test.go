package controllers_test

import (
	. "choreboard/controllers"
	. "choreboard/models"
	"choreboard/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chore Controller", func() {
	var (
		fakeController *ChoreController
		w              *httptest.ResponseRecorder
		req            *http.Request
		err            error
	)

	Describe("AllChores", func() {
		Context("when DB operation is succesful", func() {
			BeforeEach(func() {
				fakeController = NewChoreController(&fakeChoreServices{})

				req, err = http.NewRequest("GET", "http://example.com/user/guid-user-id/chores", nil)

				w = httptest.NewRecorder()
				ginContext := &gin.Context{
					Writer: w,
					Req:    req,
				}
				fakeController.AllCreatedChores(ginContext)
			})

			It("should return all Chores", func() {

				expected, _ := json.Marshal([]Chore{
					Chore{"123", "test1", "Description1", "guid-user-id", &now},
					Chore{"1234", "test2", "Description2", "guid-user-id", &now},
					Chore{"12345", "test3", "Description3", "guid-user-id", &now},
					Chore{"123456", "test4", "Description4", "guid-user-id", &now},
				})

				Expect(expected).Should(MatchJSON(w.Body))

			})

			It("should return 200 http code", func() {
				Expect(http.StatusOK).Should(Equal(w.Code))
			})
		})

		Context("when DB operation fails", func() {
			BeforeEach(func() {
				fakeController = NewChoreController(&fakeChoreErrorServices{})

				req, err = http.NewRequest("GET", "http://example.com/user/guid-user-id/chores", nil)

				w = httptest.NewRecorder()
				ginContext := &gin.Context{
					Writer: w,
					Req:    req,
				}
				fakeController.AllCreatedChores(ginContext)
			})

			It("should return error message", func() {
				expected, _ := json.Marshal(&utils.Error{500, "DB Error"})
				Expect(expected).Should(MatchJSON(w.Body))
			})

			It("should return 500 http code", func() {
				Expect(http.StatusInternalServerError).Should(Equal(w.Code))
			})
		})

	})

})

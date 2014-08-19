package main

import (
	"github.com/emicklei/go-restful"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"log"	
	"strings"
	"io/ioutil"
)

var user, updatedUser User
var service UserService
var userAsJson, updatedUserAsJson []byte
var response *restful.Response
var httpResponse *httptest.ResponseRecorder

var _ = BeforeSuite(func() {
    restful.DefaultResponseContentType(restful.MIME_JSON)
	user = User{"1", "Andy"}
	userAsJson, _ = json.Marshal(user)
	updatedUser = User{"1", "Andrew"}
	updatedUserAsJson, _ = json.Marshal(updatedUser)
    httpResponse = httptest.NewRecorder()
    response = restful.NewResponse(httpResponse)
})

func createUser(user []byte) {
    request, err := http.NewRequest("PUT", "/users/1", strings.NewReader(string(user)))
    request.Header.Set("Content-Type", "application/json")
    Ω(err).ShouldNot(HaveOccurred())

    service.createUser("1", restful.NewRequest(request), response)

}

var _ = Describe("UserService", func() {
	BeforeEach(func() {
	    service = UserService{map[string]User{}}
    })

	PDescribe("When finding users", func() {

		Context("that doesn't exist", func() {

			It("should fail", func() {
			    request, _ := http.NewRequest("GET", "/users/1", strings.NewReader(string(userAsJson)))
			    request.Header.Set("Content-Type", "application/json")

			    service.findUser("1", response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusNotFound))
			})
		})

		Context("that exist", func() {

			It("should succeed", func() {
				createUser(userAsJson)

			    request, _ := http.NewRequest("GET", "/users/1", strings.NewReader(string(userAsJson)))
			    request.Header.Set("Content-Type", "application/json")

			    service.findUser("1", response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			})
		})

	})


	Describe("When adding users", func() {
		Context("that doesn't exist", func() {

			It("should succeed", func() {
				createUser(userAsJson)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())

			    newUser := new(User)
			    json.Unmarshal(body, newUser)
			    Ω(newUser).Should(Equal(&user))
			})
		})

		Context("that already exists", func() {

			It("should succeed and overwrite the user", func() {
				createUser(userAsJson)

				createUser(updatedUserAsJson)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())
			    log.Printf("%s", body)

			})
		})
	})
})

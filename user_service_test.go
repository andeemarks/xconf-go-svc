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
var userAsJsonString, updatedUserAsJsonString string
var response *restful.Response
var httpResponse *httptest.ResponseRecorder

var _ = BeforeSuite(func() {
    restful.DefaultResponseContentType(restful.MIME_JSON)
	user = User{"1", "Andy"}
	userAsJson, _ := json.Marshal(user)
	userAsJsonString = string(userAsJson)
	updatedUser = User{"1", "Andrew"}
	updatedUserAsJson, _ := json.Marshal(updatedUser)
	updatedUserAsJsonString = string(updatedUserAsJson)
})

func createUser(user string) {
	// log.Printf("Adding: %s", user)
    request, err := http.NewRequest("PUT", "/users/1", strings.NewReader(user))
    request.Header.Set("Content-Type", "application/json")
    Ω(err).ShouldNot(HaveOccurred())

    service.createUser("1", restful.NewRequest(request), response)

}

var _ = Describe("UserService", func() {
	BeforeEach(func() {
	    service = UserService{map[string]User{}}
	    httpResponse = httptest.NewRecorder()
	    response = restful.NewResponse(httpResponse)
    })

	Describe("When finding users", func() {

		Context("that doesn't exist", func() {

			It("should fail", func() {
			    request, _ := http.NewRequest("GET", "/users/1", nil)
			    request.Header.Set("Content-Type", "application/json")

			    service.findUser("1", response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusNotFound))
			})
		})

		Context("that exist", func() {

			It("should succeed", func() {
				createUser(userAsJsonString)

			    request, _ := http.NewRequest("GET", "/users/1", nil)
			    request.Header.Set("Content-Type", "application/json")

			    service.findUser("1", response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			})
		})

	})

	Describe("when updating users", func() {
		Context("that don't exist", func() {

			It("user should be added", func() {
			    request, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(userAsJsonString))
			    request.Header.Set("Content-Type", "application/json")

			    service.UpdateUser(restful.NewRequest(request), response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusOK))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())

			    addedUser := new(User)
			    json.Unmarshal(body, addedUser)
			    Ω(addedUser).Should(Equal(&user))
			})
		})

		PContext("that do exist", func() {

			It("user should be updated", func() {
				createUser(userAsJsonString)

			    request, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(updatedUserAsJsonString))
			    request.Header.Set("Content-Type", "application/json")

			    service.UpdateUser(restful.NewRequest(request), response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())

			    addedUser := new(User)
			    json.Unmarshal(body, addedUser)
			    Ω(addedUser).Should(Equal(&updatedUser))
			})
		})

	})

	Describe("When deleting users", func() {
		Context("that don't exist", func() {
			It("should fail silently", func() {
			    request, _ := http.NewRequest("DELETE", "/users/1", nil)
			    request.Header.Set("Content-Type", "application/json")

				service.removeUser("1", restful.NewRequest(request), response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusOK))

			})
		})	

		Context("that do exist", func() {
			It("should not be possible to find those users afterwards", func() {
				createUser(userAsJsonString)

			    request, _ := http.NewRequest("DELETE", "/users/1", nil)
			    request.Header.Set("Content-Type", "application/json")

				service.removeUser("1", restful.NewRequest(request), response)

			    service.findUser("1", response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusNotFound))

			})
		})	
	})

	Describe("When adding users", func() {
		Context("that don't exist", func() {

			It("should succeed", func() {
				createUser(userAsJsonString)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())

			    addedUser := new(User)
			    json.Unmarshal(body, addedUser)
			    Ω(addedUser).Should(Equal(&user))
			})
		})

		Context("that already exists", func() {

			It("should succeed and overwrite the user", func() {
				createUser(userAsJsonString)

				createUser(updatedUserAsJsonString)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())
			    log.Printf("%s", body)

			})
		})
	})
})

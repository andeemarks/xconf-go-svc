package main

import (
	"github.com/emicklei/go-restful"
	// "encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"log"	
	"strings"
	"io/ioutil"
)

var user string
var updatedUser string
var service UserService
var userAsJson []byte
var response *restful.Response
var httpResponse *httptest.ResponseRecorder

var _ = BeforeSuite(func() {
    restful.DefaultResponseContentType(restful.MIME_JSON)
	user = `{"Id": "1", "Name": "Andy"}`
	updatedUser = `{"Id": "1", "Name": "Andrew"}`
    httpResponse = httptest.NewRecorder()
    response = restful.NewResponse(httpResponse)
	// userAsJson, _ := json.Marshal(User{"1", "Andy"})
})

func createUser(user string) {
    request, err := http.NewRequest("PUT", "http://localhost:8080/users/1", strings.NewReader(user))
    request.Header.Set("Content-Type", "application/json")
    Ω(err).ShouldNot(HaveOccurred())

    service.createUser("1", restful.NewRequest(request), response)

}

var _ = Describe("UserService", func() {
	BeforeEach(func() {
	    service = UserService{map[string]User{}}
    })

	PIt("should be Swagger compliant", func() {
	    uri := "http://localhost:8080/apidocs.json"
	    response, err := http.Get(uri)
	    Ω(err).ShouldNot(HaveOccurred())

	    body, err := ioutil.ReadAll(response.Body)
	    Ω(err).ShouldNot(HaveOccurred())
	    Ω(response.StatusCode).Should(Equal(http.StatusOK))
	    Ω(body).Should(ContainSubstring("swagger"))
	})
	
	PDescribe("When finding users", func() {

		Context("that doesn't exist", func() {

			It("should fail", func() {
			    uri := "http://localhost:8080/users/1"
			    response, err := http.Get(uri)
			    Ω(err).ShouldNot(HaveOccurred())
			    Ω(response.StatusCode).Should(Equal(http.StatusNotFound))
			})
		})
	})


	Describe("When adding users", func() {
		Context("that doesn't exist", func() {

			It("should succeed", func() {
				createUser(user)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())
			    // Ω(body).Should(Equal(user))
			    log.Printf("%s", body)
			})
		})

		Context("that already exists", func() {

			It("should succeed and overwrite the user", func() {
				createUser(user)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))

				createUser(updatedUser)

			    // request, err := http.NewRequest("PUT", "http://localhost:8080/users/1", strings.NewReader(updatedUser))
			    // request.Header.Set("Content-Type", "application/json")
			    // Ω(err).ShouldNot(HaveOccurred())

			    // service.createUser("1", restful.NewRequest(request), response)

			    Ω(response.StatusCode()).Should(Equal(http.StatusCreated))
			    body, err := ioutil.ReadAll(httpResponse.Body)
			    Ω(err).ShouldNot(HaveOccurred())
			    log.Printf("%s", body)

			})
		})
	})
})

package xconf_go_svc_test

import (
	. "github.com/andeemarks/xconf-go-svc"
	"github.com/emicklei/go-restful"
	"net/http"
	"net/http/httptest"
	"strings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {
	var (
        service  UserService
        request  *restful.Request
        response *restful.Response
    )

	BeforeEach(func() {
		service = UserService{map[string]User{}}
		restful.SetCacheReadEntity(true)
		httpWriter := httptest.NewRecorder()
		response = restful.NewResponse(httpWriter) 
    })

	Describe("Creating users", func() {
        Context("With the necessary fields", func() {
            It("should succeed", func() {
				bodyReader := strings.NewReader("{\"Id\": \"1\",\"Name\": \"Andy\"}")
				httpRequest, _ := http.NewRequest("PUT", "/users", bodyReader)
				httpRequest.Header.Set("Content-Type", "application/json") 		
				request = restful.NewRequest(httpRequest)
            	service.CreateUser(request, response)
                Expect(response.StatusCode()).To(Equal(http.StatusCreated))
            })
        })
    })

	Describe("Finding users", func() {
        Context("With a non-existant user id", func() {
            It("should return an error", func() {
            	service.FindUser(request, response)
                Expect(response.StatusCode()).To(Equal(http.StatusNotFound))
            })
        })

        PContext("With an existing user id", func() {
            It("should succeed", func() {
            	service.FindUser(request, response)
                Expect(response.StatusCode()).To(Equal(http.StatusOK))
            })
        })
    })
})

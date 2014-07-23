package xconf_go_svc_test

import (
	. "github.com/andeemarks/xconf-go-svc"
	"github.com/emicklei/go-restful"
	"net/http"
	"net/http/httptest"
	// "strings"
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
		request = new(restful.Request)
		httpWriter := httptest.NewRecorder()
		response = restful.NewResponse(httpWriter) 
    })

	Describe("Finding users", func() {
        Context("With a non-existant user id", func() {
            It("should return an error", func() {
            	service.FindUser(request, response)
                Expect(response.StatusCode()).To(Equal(http.StatusNotFound))
            })
        })

        Context("With an existing user id", func() {
            It("should return success", func() {
            	service.FindUser(request, response)
                Expect(response.StatusCode()).To(Equal(http.StatusOK))
            })
        })
    })
})

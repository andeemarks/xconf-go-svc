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
		bodyReader := strings.NewReader("<Sample><Value>42</Value></Sample>")
		httpRequest, _ := http.NewRequest("GET", "/test", bodyReader)
		request := &restful.Request{Request: httpRequest}
		httpWriter := httptest.NewRecorder()
		response := &restful.Response{httpWriter, "*/*", []string{"*/*"}, 0, 0}
    })

	Describe("Finding users", func() {
        Context("With a non-existant user id", func() {
            It("should return an error", func() {
            	service.FindUser(request, response)
                Expect(response).To(Equal("NOVEL"))
            })
        })

        Context("With an existing user id", func() {
            It("should return success", func() {
                // Expect(service.FindUser(request, response)).To(Equal("SHORT STORY"))
            })
        })
    })
})

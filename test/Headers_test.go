package test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/taxjar/taxjar-go"
)

var _ = Describe("Request headers:", func() {
	RegisterFailHandler(Fail)
	defer GinkgoRecover()

	var server *ghttp.Server
	var client taxjar.Config

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = taxjar.NewClient(taxjar.Config{
			APIKey: "test123",
			APIURL: server.URL(),
		})
	})

	It("should include all appropriate headers", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/v2/categories"),
			ghttp.VerifyContentType("application/json"),
			ghttp.VerifyHeader(http.Header{
				"Authorization": []string{"Bearer test123"},
			}),
		))
		client.Categories()

		userAgent := server.ReceivedRequests()[0].Header["User-Agent"][0]
		Expect(userAgent).To(MatchRegexp(`^TaxJar/Go \(.*\) taxjar-go/\d+\.\d+\.\d+$`))
	})
})

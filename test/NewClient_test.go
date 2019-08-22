package test

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/taxjar/taxjar-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "taxjar")
}

var _ = Describe("NewClient without config", func() {
	var client taxjar.Config
	BeforeEach(func() {
		client = taxjar.NewClient()
	})

	It("should allow setting APIKey", func() {
		client.APIKey = "test123"
		Expect(client.APIKey).To(Equal("test123"))
	})

	It("should allow setting APIURL", func() {
		mockURL := "https://api.mock.taxjar.com"
		client.APIURL = mockURL
		Expect(client.APIURL).To(Equal(mockURL))
	})

	It("should allow setting APIVersion", func() {
		APIVersion := "v2"
		client.APIVersion = APIVersion
		Expect(client.APIVersion).To(Equal(APIVersion))
	})

	It("should allow setting customer headers", func() {
		headers := map[string]interface{}{
			"X-TJ-Expected-Response": 422,
		}
		client.Headers = headers
		Expect(client.Headers).To(Equal(headers))
	})

	It("should exit when no APIKey is set", func() {
		fakeLogFatal := monkey.Patch(log.Fatal, func(msg ...interface{}) {
			Expect(fmt.Sprintf("%v", msg[0])).To(MatchRegexp("taxjar:"))
			panic("fake log.Fatal called")
		})
		defer fakeLogFatal.Unpatch()
		defer func() {
			if r := recover(); r == nil {
				Fail("Expected log.Fatal to be called when no APIKey is set")
			}
		}()
		client.Categories()
	})

	It("should allow setting a custom timeout", func() {
		client.Timeout = 10 * time.Minute
		Expect(client.Timeout).To(Equal(10 * time.Minute))
	})

	It("should allow setting a custom transport", func() {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: 20 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   20 * time.Second,
			ExpectContinueTimeout: 8 * time.Second,
			ResponseHeaderTimeout: 6 * time.Second,
		}
		client.Transport = transport
		Expect(client.Transport).To(Equal(transport))
	})

})

package test

import (
	"net"
	"net/http"
	"time"

	"github.com/taxjar/taxjar-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewClient with config", func() {

	It("should allow setting APIKey", func() {
		client := taxjar.NewClient(taxjar.Config{
			APIKey: "test123",
		})
		Expect(client.APIKey).To(Equal("test123"))
	})

	It("should allow setting APIURL", func() {
		mockURL := "https://api.mock.taxjar.com"
		client := taxjar.NewClient(taxjar.Config{
			APIURL: mockURL,
		})
		Expect(client.APIURL).To(Equal(mockURL))
	})

	It("should allow setting APIVersion", func() {
		APIVersion := "v2"
		client := taxjar.NewClient(taxjar.Config{
			APIVersion: APIVersion,
		})
		Expect(client.APIVersion).To(Equal(APIVersion))
	})

	It("should allow setting custom headers", func() {
		headers := map[string]interface{}{
			"X-TJ-Expected-Response": 422,
		}
		client := taxjar.NewClient(taxjar.Config{
			Headers: headers,
		})
		Expect(client.Headers).To(Equal(headers))
	})

	It("should allow setting a custom timeout", func() {
		client := taxjar.NewClient(taxjar.Config{
			Timeout: 10 * time.Minute,
		})
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
		client := taxjar.NewClient(taxjar.Config{
			Transport: transport,
		})
		Expect(client.Transport).To(Equal(transport))
	})

	It("should allow setting a custom *http.Client", func() {
		httpClient := &http.Client{
			Timeout: 10 * time.Minute,
		}
		client := taxjar.NewClient(taxjar.Config{
			HTTPClient: httpClient,
		})
		Expect(client.HTTPClient).To(Equal(httpClient))
	})

})

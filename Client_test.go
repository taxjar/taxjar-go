package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	taxjar "github.com/taxjar/taxjar-go"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestClientSetApiKey(t *testing.T) {
	client := taxjar.NewClient()
	client.APIKey = "test123"
	assert.Equal(t, "test123", client.APIKey)
}

func TestClientSetApiUrl(t *testing.T) {
	client := taxjar.NewClient()
	client.APIURL = "https://api.mock.taxjar.com"
	assert.Equal(t, "https://api.mock.taxjar.com", client.APIURL)
}

func TestClientSetApiVersion(t *testing.T) {
	client := taxjar.NewClient()
	client.APIVersion = "v2"
	assert.Equal(t, "v2", client.APIVersion)
}

func TestClientSetCustomHeaders(t *testing.T) {
	headers := map[string]interface{}{
		"X-TJ-Expected-Response": 422,
	}
	client := taxjar.NewClient()
	client.Headers = headers
	assert.Equal(t, headers, client.Headers)
}

func TestClientSetCustomTimeout(t *testing.T) {
	timeout := 10 * time.Minute
	client := taxjar.NewClient()
	client.Timeout = timeout
	assert.Equal(t, timeout, client.Timeout)
}

func TestClientSetCustomTransport(t *testing.T) {
	client := taxjar.NewClient()
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
	assert.Equal(t, transport, client.Transport)
}

func TestClientSetCustomHttpClient(t *testing.T) {
	client := taxjar.NewClient()
	httpClient := &http.Client{
		Timeout: 10 * time.Minute,
	}
	client.HTTPClient = httpClient
	assert.Equal(t, httpClient, client.HTTPClient)
}

func TestClientIncludesAppropriateHeaders(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, "Bearer mock-api-key", req.Header.Get("Authorization"))
		assert.Regexp(t, `^TaxJar/Go \(.*\) taxjar-go/\d+\.\d+\.\d+$`, req.Header["User-Agent"][0])
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.CategoriesJSON))),
			Header:     make(http.Header),
		}
	})

	client.Categories()
}

func TestClientShouldNotAllowOverwritingStandardHeaders(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.NotEqual(t, "a", req.Header.Get("Authorization"))
		assert.NotEqual(t, "b", req.Header.Get("Content-Type"))
		assert.NotEqual(t, "c", req.Header.Get("user-agent"))
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.CategoriesJSON))),
			Header:     make(http.Header),
		}
	})
	badHeaders := map[string]interface{}{
		// check different upper/lower casings
		"Authorization": "a",
		"Content-type":  "b",
		"user-agent":    "c",
	}
	client.Headers = badHeaders

	client.Categories()
}

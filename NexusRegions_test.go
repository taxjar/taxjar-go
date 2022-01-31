package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestNexusRegionsOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/nexus/regions", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.NexusRegionsJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.NexusRegions()

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.NexusRegions, *res)
}

func TestNexusRegionsOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/nexus/regions", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.NexusRegions()

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

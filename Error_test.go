package taxjar_test

import (
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestErrorOutput(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.Categories()

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

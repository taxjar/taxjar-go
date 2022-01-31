package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestDeleteOrderOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders/123", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.DeleteOrderJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.DeleteOrder("123")

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.DeleteOrder, *res)
}

func TestDeleteOrderOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders/123", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.DeleteOrder("123")

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

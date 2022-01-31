package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestDeleteCustomerOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers/123", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.DeleteCustomerJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.DeleteCustomer("123")

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.DeleteCustomer, *res)
}

func TestDeleteCustomerOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers/123", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.DeleteCustomer("123")

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

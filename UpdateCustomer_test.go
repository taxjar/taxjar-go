package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	taxjar "github.com/taxjar/taxjar-go"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestUpdateCustomerOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers/123", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.UpdateCustomerJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.UpdateCustomer(taxjar.UpdateCustomerParams{
		CustomerID:    "123",
		ExemptionType: "non_exempt",
		Name:          "Initech",
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.UpdateCustomer, *res)
}

func TestUpdateCustomerOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers/123", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.UpdateCustomer(taxjar.UpdateCustomerParams{
		CustomerID:    "123",
		ExemptionType: "non_exempt",
		Name:          "Initech",
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

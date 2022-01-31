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

func TestCreateCustomerOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.CreateCustomerJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.CreateCustomer(taxjar.CreateCustomerParams{
		CustomerID:    "123",
		ExemptionType: "wholesale",
		Name:          "Initech",
		ExemptRegions: []taxjar.ExemptRegion{
			{
				Country: "US",
				State:   "TX",
			},
		},
		Country: "US",
		State:   "TX",
		Zip:     "78744",
		City:    "Austin",
		Street:  "4120 Freidrich Lane",
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.CreateCustomer, *res)
}

func TestCreateCustomerOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.CreateCustomer(taxjar.CreateCustomerParams{
		CustomerID:    "123",
		ExemptionType: "wholesale",
		Name:          "Initech",
		ExemptRegions: []taxjar.ExemptRegion{
			{
				Country: "US",
				State:   "TX",
			},
		},
		Country: "US",
		State:   "TX",
		Zip:     "78744",
		City:    "Austin",
		Street:  "4120 Freidrich Lane",
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

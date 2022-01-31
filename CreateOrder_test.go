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

func TestCreateOrderOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.CreateOrderJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.CreateOrder(taxjar.CreateOrderParams{
		TransactionID:   "24",
		TransactionDate: "2019/08/26",
		Provider:        "api",
		FromCountry:     "US",
		FromZip:         "94043",
		FromState:       "CA",
		FromCity:        "Mountain View",
		FromStreet:      "311 Moffett Blvd",
		ToCountry:       "US",
		ToZip:           "10019",
		ToState:         "NY",
		ToCity:          "New York",
		ToStreet:        "1697 Broadway",
		Amount:          50,
		Shipping:        5,
		SalesTax:        0,
		CustomerID:      "123",
		ExemptionType:   "non_exempt",
		LineItems: []taxjar.OrderLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "10-12345-987",
				Description:       "10-gallon Hat",
				ProductTaxCode:    "20010",
				UnitPrice:         45,
				Discount:          0,
				SalesTax:          0,
			},
		},
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.CreateOrder, *res)
}

func TestCreateOrderOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.CreateOrder(taxjar.CreateOrderParams{
		TransactionID:   "24",
		TransactionDate: "2019/08/26",
		Provider:        "api",
		FromCountry:     "US",
		FromZip:         "94043",
		FromState:       "CA",
		FromCity:        "Mountain View",
		FromStreet:      "311 Moffett Blvd",
		ToCountry:       "US",
		ToZip:           "10019",
		ToState:         "NY",
		ToCity:          "New York",
		ToStreet:        "1697 Broadway",
		Amount:          50,
		Shipping:        5,
		SalesTax:        0,
		CustomerID:      "123",
		ExemptionType:   "non_exempt",
		LineItems: []taxjar.OrderLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "10-12345-987",
				Description:       "10-gallon Hat",
				ProductTaxCode:    "20010",
				UnitPrice:         45,
				Discount:          0,
				SalesTax:          0,
			},
		},
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

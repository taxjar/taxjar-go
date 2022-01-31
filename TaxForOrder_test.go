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

func TestTaxForOrderOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/taxes", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.USTaxForOrderJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
		FromCountry:   "US",
		FromZip:       "92093",
		FromState:     "CA",
		FromCity:      "La Jolla",
		FromStreet:    "9500 Gilman Drive",
		ToCountry:     "US",
		ToZip:         "90002",
		ToState:       "CA",
		ToCity:        "Los Angeles",
		ToStreet:      "1335 E 103rd St",
		Amount:        15,
		Shipping:      1.5,
		CustomerID:    "123",
		ExemptionType: "non_exempt",
		LineItems: []taxjar.TaxLineItem{
			{
				ID:             "1",
				Quantity:       1,
				ProductTaxCode: "20010",
				UnitPrice:      15,
				Discount:       0,
			},
		},
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.USTaxForOrder, *res)
}

func TestTaxForOrderOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/taxes", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
		FromCountry:   "US",
		FromZip:       "92093",
		FromState:     "CA",
		FromCity:      "La Jolla",
		FromStreet:    "9500 Gilman Drive",
		ToCountry:     "US",
		ToZip:         "90002",
		ToState:       "CA",
		ToCity:        "Los Angeles",
		ToStreet:      "1335 E 103rd St",
		Amount:        15,
		Shipping:      1.5,
		CustomerID:    "123",
		ExemptionType: "non_exempt",
		LineItems: []taxjar.TaxLineItem{
			{
				ID:             "1",
				Quantity:       1,
				ProductTaxCode: "20010",
				UnitPrice:      15,
				Discount:       0,
			},
		},
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

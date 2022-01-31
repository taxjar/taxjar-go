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

func TestCreateRefundOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/refunds", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.CreateRefundJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.CreateRefund(taxjar.CreateRefundParams{
		TransactionID:          "24-refund",
		TransactionReferenceID: "24",
		TransactionDate:        "2019/08/26",
		Provider:               "api",
		FromCountry:            "US",
		FromZip:                "94043",
		FromState:              "CA",
		FromCity:               "Mountain View",
		FromStreet:             "311 Moffett Blvd",
		ToCountry:              "US",
		ToZip:                  "10019",
		ToState:                "NY",
		ToCity:                 "New York",
		ToStreet:               "1697 Broadway",
		Amount:                 -111,
		Shipping:               -0,
		SalesTax:               -10.3,
		CustomerID:             "123",
		ExemptionType:          "non_exempt",
		LineItems: []taxjar.RefundLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "10-12345-987",
				Description:       "10-gallon Hat",
				ProductTaxCode:    "20010",
				UnitPrice:         -0,
				Discount:          -0,
				SalesTax:          -0,
			},
			{
				ID:                "2",
				Quantity:          1,
				ProductIdentifier: "78-95432-101",
				Description:       "Extra-long Chaps",
				ProductTaxCode:    "20010",
				UnitPrice:         -111,
				Discount:          -0,
				SalesTax:          -9.85,
			},
		},
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.CreateRefund, *res)
}

func TestCreateRefundOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/refunds", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.CreateRefund(taxjar.CreateRefundParams{
		TransactionID:          "24-refund",
		TransactionReferenceID: "24",
		TransactionDate:        "2019/08/26",
		Provider:               "api",
		FromCountry:            "US",
		FromZip:                "94043",
		FromState:              "CA",
		FromCity:               "Mountain View",
		FromStreet:             "311 Moffett Blvd",
		ToCountry:              "US",
		ToZip:                  "10019",
		ToState:                "NY",
		ToCity:                 "New York",
		ToStreet:               "1697 Broadway",
		Amount:                 -111,
		Shipping:               -0,
		SalesTax:               -10.3,
		CustomerID:             "123",
		ExemptionType:          "non_exempt",
		LineItems: []taxjar.RefundLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "10-12345-987",
				Description:       "10-gallon Hat",
				ProductTaxCode:    "20010",
				UnitPrice:         -0,
				Discount:          -0,
				SalesTax:          -0,
			},
			{
				ID:                "2",
				Quantity:          1,
				ProductIdentifier: "78-95432-101",
				Description:       "Extra-long Chaps",
				ProductTaxCode:    "20010",
				UnitPrice:         -111,
				Discount:          -0,
				SalesTax:          -9.85,
			},
		},
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

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

func TestUpdateOrderOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders/24", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.UpdateOrderJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.UpdateOrder(taxjar.UpdateOrderParams{
		TransactionID: "24",
		Amount:        161,
		SalesTax:      10.3,
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
			{
				ID:                "2",
				Quantity:          1,
				ProductIdentifier: "78-95432-101",
				Description:       "Extra-long Chaps",
				ProductTaxCode:    "20010",
				UnitPrice:         111,
				Discount:          0,
				SalesTax:          9.85,
			},
		},
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.UpdateOrder, *res)
}

func TestUpdateOrderOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders/24", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.UpdateOrder(taxjar.UpdateOrderParams{
		TransactionID: "24",
		Amount:        161,
		SalesTax:      10.3,
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
			{
				ID:                "2",
				Quantity:          1,
				ProductIdentifier: "78-95432-101",
				Description:       "Extra-long Chaps",
				ProductTaxCode:    "20010",
				UnitPrice:         111,
				Discount:          0,
				SalesTax:          9.85,
			},
		},
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

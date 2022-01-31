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

func TestUpdateRefundOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/refunds/24-refund", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.UpdateRefundJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.UpdateRefund(taxjar.UpdateRefundParams{
		TransactionID:          "24-refund",
		TransactionReferenceID: "24",
		Shipping:               -5,
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.UpdateRefund, *res)
}

func TestUpdateRefundOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/refunds/24-refund", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.UpdateRefund(taxjar.UpdateRefundParams{
		TransactionID:          "24-refund",
		TransactionReferenceID: "24",
		Shipping:               -5,
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

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

func TestListOrdersOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders?transaction_date=2019%2F08%2F26", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.ListOrdersJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.ListOrders(taxjar.ListOrdersParams{
		TransactionDate: "2019/08/26",
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.ListOrders, *res)
}

func TestListOrdersOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/transactions/orders?transaction_date=2019%2F08%2F26", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.ListOrders(taxjar.ListOrdersParams{
		TransactionDate: "2019/08/26",
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

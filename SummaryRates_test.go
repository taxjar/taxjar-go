package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestSummaryRatesOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/summary_rates", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.SummaryRatesJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.SummaryRates()

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.SummaryRates, *res)
}

func TestSummaryRatesOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/summary_rates", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.SummaryRates()

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

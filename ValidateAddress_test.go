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

func TestValidateAddressOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/addresses/validate", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.ValidateAddressJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.ValidateAddress(taxjar.ValidateAddressParams{
		Country: "US",
		State:   "AZ",
		Zip:     "85297",
		City:    "Gilbert",
		Street:  "3301 South Greenfield Rd",
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.ValidateAddress, *res)
}

func TestValidateAddressOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/addresses/validate", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.ValidateAddress(taxjar.ValidateAddressParams{
		Country: "US",
		State:   "AZ",
		Zip:     "85297",
		City:    "Gilbert",
		Street:  "3301 South Greenfield Rd",
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestCategoriesOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/categories", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.CategoriesJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.Categories()

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, mock.Categories, *res)
}

func TestCategoriesOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/categories", req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Header:     make(http.Header),
		}
	})

	res, err := client.Categories()

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

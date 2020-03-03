package taxjar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// DefaultTimeout for requests is `30 * time.Second`․
//
// Override this by setting `Config.Timeout` to a different time value․
const DefaultTimeout = 30 * time.Second

// DefaultTransport is the default `*http.Transport` for requests․
//
// Override this by setting `Config.Transport` to a different `*http.Transport` (from net/http package - https://godoc.org/net/http#Transport)․
var DefaultTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
	}).DialContext,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 4 * time.Second,
	ResponseHeaderTimeout: 3 * time.Second,
}

func (client *Config) setTimeouts() {
	if client.HTTPClient == nil {
		client.HTTPClient = &http.Client{}
	}
	if client.Timeout > 0 {
		client.HTTPClient.Timeout = client.Timeout
	} else if client.HTTPClient.Timeout <= 0 {
		client.HTTPClient.Timeout = DefaultTimeout
	}
	if client.Transport != nil {
		client.HTTPClient.Transport = client.Transport
	} else if client.HTTPClient.Transport == nil {
		client.HTTPClient.Transport = DefaultTransport
	}
}

func (client *Config) url(endpoint string) string {
	return fmt.Sprintf("%v/%v/%v", client.APIURL, client.APIVersion, endpoint)
}

func (client *Config) addHeaders(req *http.Request) {
	if client.APIKey == "" {
		log.Fatal("taxjar: missing `APIKey` field must be set on client")
	}
	req.Header.Add("Authorization", "Bearer "+client.APIKey)
	req.Header.Add("Content-type", "application/json")
	for key, val := range client.Headers {
		val, _ := val.(string)
		if key != "Authorization" && key != "Content-type" {
			req.Header.Add(key, val)
		}
	}
}

func addQueryParams(req *http.Request, params interface{}) error {
	queryParams, err := query.Values(params)
	if err != nil {
		return err
	}
	req.URL.RawQuery = queryParams.Encode()
	return nil
}

func (client *Config) sendRequest(req *http.Request) ([]byte, error) {
	client.addHeaders(req)
	client.setTimeouts()
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		err := new(Error)
		json.Unmarshal(body, err)
		return nil, err
	}
	return body, nil
}

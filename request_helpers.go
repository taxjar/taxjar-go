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

// DefaultTimeout - 30 * time.Second
const DefaultTimeout = 30 * time.Second

/*
DefaultTransport - &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
	}).DialContext,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 4 * time.Second,
	ResponseHeaderTimeout: 3 * time.Second,
}
*/
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
	if client.Timeout > 0 {
		client.httpClient.Timeout = client.Timeout
	} else {
		client.httpClient.Timeout = DefaultTimeout
	}
	if client.Transport != nil {
		client.httpClient.Transport = client.Transport
	} else {
		client.httpClient.Transport = DefaultTransport
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

func addQueryParams(req *http.Request, params interface{}) {
	queryParams, _ := query.Values(params)
	req.URL.RawQuery = queryParams.Encode()
}

func (client *Config) sendRequest(req *http.Request) ([]byte, error) {
	client.addHeaders(req)
	client.setTimeouts()
	res, err := client.httpClient.Do(req)
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

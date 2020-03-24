package taxjar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
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

func getUserAgent() string {
	platform := runtime.GOOS
	uname, err := exec.LookPath("uname")
	if err == nil {
		cmd := exec.Command(uname, "-a")
		var out bytes.Buffer
		cmd.Stderr = nil
		cmd.Stdout = &out
		err = cmd.Run()
		if err == nil {
			platform = strings.TrimSpace(out.String())
		}
	}

	goVersion := runtime.Version()
	re, _ := regexp.Compile(`go(.+)`)
	matches := re.FindStringSubmatch(goVersion)
	if len(matches) > 1 {
		goVersion = fmt.Sprintf("%v", matches[1])
	}

	return fmt.Sprintf("TaxJar/Go (%v; %v; go %v) taxjar-go/%v", platform, runtime.GOARCH, goVersion, version)
}

func (client *Config) addHeaders(req *http.Request) {
	if client.APIKey == "" {
		log.Fatal("taxjar: missing `APIKey` field must be set on client")
	}

	req.Header.Add("Authorization", "Bearer "+client.APIKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", getUserAgent())

	for key, val := range client.Headers {
		if key = http.CanonicalHeaderKey(key); key != "Authorization" && key != "Content-Type" && key != "User-Agent" {
			req.Header.Add(key, val.(string))
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

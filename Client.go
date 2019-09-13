// Package taxjar - Official Go API client from TaxJar (https://www.taxjar.com)
//
// For more information, see our:
//
// • Go Quickstart Guide (https://developers.taxjar.com/api/guides/go/)
//
// • API docs (https://developers.taxjar.com/api/reference/?go)
//
// • API Guides (https://developers.taxjar.com/api/guides)
//
// • Integration Guides (https://developers.taxjar.com/integrations)
//
// • README (https://github.com/taxjar/taxjar-go/blob/master/README.md)
package taxjar

import (
	"net/http"
	"time"
)

// Config is the structure for configuring a `taxjar` client․ Pass a `Config` to `NewClient` to instantiate a client․
//
// See below for default values․
type Config struct {
	APIURL     string // default: "https://api.taxjar.com"
	APIKey     string
	APIVersion string // default: "v2"
	Headers    map[string]interface{}
	HTTPClient *http.Client    // default: `&http.Client{}` (from net/http package - https://godoc.org/net/http#Client)
	Timeout    time.Duration   // default: `30 * time.Second`
	Transport  *http.Transport /* default:
	&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 4 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
	}*/
}

// DefaultAPIURL ("https://api.taxjar.com") is the default base API URL for each request to TaxJar․
//
// Override this by setting `Config.APIURL` to a different string․
const DefaultAPIURL = "https://api.taxjar.com"

// SandboxAPIURL ("https://api.sandbox.taxjar.com") is the base API URL to send requests to TaxJar's sandbox environment․
//
// Use by setting `Config.APIURL` to `taxjar.SandboxAPIURL`․
//
// Please note that TaxJar's sandbox environment requires a TaxJar Plus account (https://www.taxjar.com/plus/)․
//
// See https://developers.taxjar.com/api/reference/?go#sandbox-environment for more details․
const SandboxAPIURL = "https://api.sandbox.taxjar.com"

// DefaultAPIVersion ("v2") is the default TaxJar API version․
//
// Override this by setting `Config.APIVersion` to a different string․
const DefaultAPIVersion = "v2"

// NewClient instantiates a new `taxjar` client․
//
// Configure the client by passing a `Config` to `NewClient` or by setting configuration values such as `APIURL`, `APIKey`, `APIVersion`, `Headers`, `HTTPClient`, `Timeout`, and `Transport` after instantiation․
//
// NewClient returns a client (type `Config`), on which you can call other methods to interact with TaxJar's API such as `Categories`, `TaxForOrder`, `CreateOrder`, etc․
//
// See our Go Quickstart Guide for more usage details and background: https://developers.taxjar.com/api/guides/go/#go-quickstart
func NewClient(config ...Config) Config {
	var _config Config
	if len(config) < 1 {
		_config = Config{}
	} else {
		_config = config[0]
	}
	if _config.APIURL == "" {
		_config.APIURL = DefaultAPIURL
	}
	if _config.APIVersion == "" {
		_config.APIVersion = DefaultAPIVersion
	}
	return _config
}

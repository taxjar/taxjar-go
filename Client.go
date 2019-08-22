// Package taxjar - Official Go API client from TaxJar
package taxjar

import (
	"net/http"
	"time"
)

// Config - TODO (document this)
type Config struct {
	APIURL     string
	APIKey     string
	APIVersion string
	Headers    map[string]interface{}
	httpClient *http.Client
	Timeout    time.Duration
	Transport  *http.Transport
}

// DefaultAPIURL - TODO (document this)
const DefaultAPIURL = "https://api.taxjar.com"

// SandboxAPIURL - TODO (document this)
const SandboxAPIURL = "https://api.sandbox.taxjar.com"

// DefaultAPIVersion - TODO (document this)
const DefaultAPIVersion = "v2"

// NewClient - TODO (document this)
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
	_config.httpClient = &http.Client{}
	return _config
}

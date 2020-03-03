package taxjar

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (client *Config) get(endpoint string, params interface{}) (interface{}, error) {
	req, _ := http.NewRequest("GET", client.url(endpoint), nil)
	if params != nil {
		if err := addQueryParams(req, params); err != nil {
			return nil, err
		}
	}
	return client.sendRequest(req)
}

func (client *Config) post(endpoint string, params interface{}) (interface{}, error) {
	jsonParams, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", client.url(endpoint), bytes.NewBuffer(jsonParams))
	return client.sendRequest(req)
}

func (client *Config) put(endpoint string, params interface{}) (interface{}, error) {
	jsonParams, _ := json.Marshal(params)
	req, _ := http.NewRequest("PUT", client.url(endpoint), bytes.NewBuffer(jsonParams))
	return client.sendRequest(req)
}

func (client *Config) delete(endpoint string, params ...interface{}) (interface{}, error) {
	req, _ := http.NewRequest("DELETE", client.url(endpoint), nil)
	addQueryParams(req, params)
	return client.sendRequest(req)
}

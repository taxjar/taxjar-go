package taxjar

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (client *Config) get(endpoint string, params interface{}) (interface{}, error) {
	req, err := http.NewRequest("GET", client.url(endpoint), nil)
	if err != nil {
		return nil, err
	}
	if params != nil {
		if err := addQueryParams(req, params); err != nil {
			return nil, err
		}
	}
	return client.sendRequest(req)
}

func (client *Config) post(endpoint string, params interface{}) (interface{}, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.url(endpoint), bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, err
	}
	return client.sendRequest(req)
}

func (client *Config) put(endpoint string, params interface{}) (interface{}, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", client.url(endpoint), bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, err
	}
	return client.sendRequest(req)
}

func (client *Config) delete(endpoint string, params interface{}) (interface{}, error) {
	req, err := http.NewRequest("DELETE", client.url(endpoint), nil)
	if err != nil {
		return nil, err
	}
	if params != nil {
		if err := addQueryParams(req, params); err != nil {
			return nil, err
		}
	}
	return client.sendRequest(req)
}

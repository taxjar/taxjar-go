package taxjar

import "encoding/json"

// ValidateAddressParams - TODO (document this)
type ValidateAddressParams struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
}

// Address - TODO (document this)
type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

// ValidateAddressResponse - TODO (document this)
type ValidateAddressResponse struct {
	Addresses []Address `json:"addresses"`
}

// ValidateAddress - TODO (document this)
func (client *Config) ValidateAddress(params ValidateAddressParams) (*ValidateAddressResponse, error) {
	res, err := client.post("addresses/validate", params)
	if err != nil {
		return nil, err
	}
	addresses := new(ValidateAddressResponse)
	json.Unmarshal(res.([]byte), &addresses)
	return addresses, nil
}

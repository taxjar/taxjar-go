package taxjar

import "encoding/json"

// ValidateAddressParams should be passed to `ValidateAddress` to validate an address․
type ValidateAddressParams struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
}

// Address is the structure for an address returned from `ValidateAddress` within `ValidateAddressResponse.Addresses`․
type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

// ValidateAddressResponse is the structure returned from `ValidateAddress․`
//
// Access the returned list of address matches with `ValidateAddressResponse.Addresses`․
type ValidateAddressResponse struct {
	Addresses []Address `json:"addresses"`
}

// ValidateAddress validates a customer address and returns back a collection of address matches․
//
// Address validation requires a TaxJar Plus subscription (https://www.taxjar.com/plus/)․
//
// See https://developers.taxjar.com/api/reference/?go#post-validate-an-address for more details․
func (client *Config) ValidateAddress(params ValidateAddressParams) (*ValidateAddressResponse, error) {
	res, err := client.post("addresses/validate", params)
	if err != nil {
		return nil, err
	}

	addresses := new(ValidateAddressResponse)
	if err := json.Unmarshal(res.([]byte), &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

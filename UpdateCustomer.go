package taxjar

import "encoding/json"

// UpdateCustomerParams should be passed to `UpdateCustomer` to update an existing customer․
type UpdateCustomerParams struct {
	CustomerID    string         `json:"customer_id,omitempty"`
	ExemptionType string         `json:"exemption_type,omitempty"`
	Name          string         `json:"name,omitempty"`
	ExemptRegions []ExemptRegion `json:"exempt_regions,omitempty"`
	Country       string         `json:"country,omitempty"`
	State         string         `json:"state,omitempty"`
	Zip           string         `json:"zip,omitempty"`
	City          string         `json:"city,omitempty"`
	Street        string         `json:"street,omitempty"`
}

// UpdateCustomerResponse is the structure returned from `UpdateCustomer`․
//
// Access the updated customer with `UpdateCustomerResponse.Customer`.
type UpdateCustomerResponse struct {
	CreateCustomerResponse
}

// UpdateCustomer updates an existing customer in TaxJar․
//
// Use the updated customer's `CustomerID` when calculating tax with TaxForOrder or when creating or updating transactions․
//
// See https://developers.taxjar.com/api/reference/?go#put-update-a-customer for more details․
func (client *Config) UpdateCustomer(params UpdateCustomerParams) (*UpdateCustomerResponse, error) {
	res, err := client.put("customers/"+params.CustomerID, params)
	if err != nil {
		return nil, err
	}
	customer := new(UpdateCustomerResponse)
	if err := json.Unmarshal(res.([]byte), &customer); err != nil {
		return nil, err
	}
	return customer, nil
}

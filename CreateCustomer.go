package taxjar

import "encoding/json"

// ExemptRegion is the structure for a customer exempt region passed within `CreateCustomerParams.ExemptRegions` and returned in `CreateCustomerResponse.Customer.ExemptRegions`․
type ExemptRegion struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

// CreateCustomerParams should be passed to `CreateCustomer` to create a customer․
type CreateCustomerParams struct {
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

// Customer is the structure for a customer returned within `CreateCustomerResponse`, `UpdateCustomerResponse`, `ShowCustomerResponse`, and `DeleteCustomerResponse`․
type Customer struct {
	CustomerID    string         `json:"customer_id"`
	ExemptionType string         `json:"exemption_type"`
	ExemptRegions []ExemptRegion `json:"exempt_regions"`
	Name          string         `json:"name"`
	Country       string         `json:"country"`
	State         string         `json:"state"`
	Zip           string         `json:"zip"`
	City          string         `json:"city"`
	Street        string         `json:"street"`
}

// CreateCustomerResponse is the structure returned from `CreateCustomer`․
//
// Access the created customer with `CreateCustomerResponse.Customer`․
type CreateCustomerResponse struct {
	Customer Customer `json:"customer"`
}

// CreateCustomer creates a new customer in TaxJar․ Use the newly created customer's `CustomerID` when calculating tax with `TaxForOrder` or when creating or updating transactions․
//
// See https://developers.taxjar.com/api/reference/?go#post-create-a-customer for more details․
func (client *Config) CreateCustomer(params CreateCustomerParams) (*CreateCustomerResponse, error) {
	res, err := client.post("customers", params)
	if err != nil {
		return nil, err
	}
	customer := new(CreateCustomerResponse)
	json.Unmarshal(res.([]byte), &customer)
	return customer, nil
}

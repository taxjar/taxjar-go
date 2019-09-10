package taxjar

import "encoding/json"

// ExemptRegion - TODO (document this)
type ExemptRegion struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

// CreateCustomerParams - TODO (document this)
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

// Customer - TODO (document this)
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

// CreateCustomerResponse - TODO (document this)
type CreateCustomerResponse struct {
	Customer Customer `json:"customer"`
}

// CreateCustomer - TODO (document this)
func (client *Config) CreateCustomer(params CreateCustomerParams) (*CreateCustomerResponse, error) {
	res, err := client.post("customers", params)
	if err != nil {
		return nil, err
	}
	customer := new(CreateCustomerResponse)
	json.Unmarshal(res.([]byte), &customer)
	return customer, nil
}

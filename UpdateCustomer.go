package taxjar

import "encoding/json"

// UpdateCustomerParams - TODO (document this)
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

// UpdateCustomerResponse - TODO (document this)
type UpdateCustomerResponse struct {
	CreateCustomerResponse
}

// UpdateCustomer - TODO (document this)
func (client *Config) UpdateCustomer(params UpdateCustomerParams) (*UpdateCustomerResponse, error) {
	res, err := client.put("customers/"+params.CustomerID, params)
	if err != nil {
		return nil, err
	}
	customer := new(UpdateCustomerResponse)
	json.Unmarshal(res.([]byte), &customer)
	return customer, nil
}

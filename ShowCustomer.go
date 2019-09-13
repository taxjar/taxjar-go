package taxjar

import "encoding/json"

// ShowCustomerResponse - TODO (document this)
type ShowCustomerResponse struct {
	CreateCustomerResponse
}

// ShowCustomer - TODO (document this)
func (client *Config) ShowCustomer(customerID string) (*ShowCustomerResponse, error) {
	res, err := client.get("customers/" + customerID)
	if err != nil {
		return nil, err
	}
	customer := new(ShowCustomerResponse)
	json.Unmarshal(res.([]byte), &customer)
	return customer, nil
}

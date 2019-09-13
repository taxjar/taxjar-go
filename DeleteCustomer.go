package taxjar

import "encoding/json"

// DeleteCustomerResponse - TODO (document this)
type DeleteCustomerResponse struct {
	CreateCustomerResponse
}

// DeleteCustomer - TODO (document this)
func (client *Config) DeleteCustomer(customerID string) (*DeleteCustomerResponse, error) {
	res, err := client.delete("customers/" + customerID)
	if err != nil {
		return nil, err
	}
	customer := new(DeleteCustomerResponse)
	json.Unmarshal(res.([]byte), &customer)
	return customer, nil
}

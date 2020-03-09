package taxjar

import "encoding/json"

// DeleteCustomerResponse is the structure returned from `DeleteCustomer`․
//
// Access the deleted customer with `DeleteCustomerResponse.Customer`․
type DeleteCustomerResponse struct {
	CreateCustomerResponse
}

// DeleteCustomer deletes a customer in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#delete-delete-a-customer for more details․
func (client *Config) DeleteCustomer(customerID string) (*DeleteCustomerResponse, error) {
	res, err := client.delete("customers/"+customerID, nil)
	if err != nil {
		return nil, err
	}
	customer := new(DeleteCustomerResponse)
	if err := json.Unmarshal(res.([]byte), &customer); err != nil {
		return nil, err
	}
	return customer, nil
}

package taxjar

import "encoding/json"

// ShowCustomerResponse is the structure returned from `ShowCustomer`․
//
// Access the customer with `ShowCustomerResponse.Customer`․
type ShowCustomerResponse struct {
	CreateCustomerResponse
}

// ShowCustomer shows an existing customer in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#get-show-a-customer for more details․
func (client *Config) ShowCustomer(customerID string) (*ShowCustomerResponse, error) {
	res, err := client.get("customers/"+customerID, nil)
	if err != nil {
		return nil, err
	}

	customer := new(ShowCustomerResponse)
	if err := json.Unmarshal(res.([]byte), &customer); err != nil {
		return nil, err
	}

	return customer, nil
}

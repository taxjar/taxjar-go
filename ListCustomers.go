package taxjar

import "encoding/json"

// ListCustomersResponse is the structure returned from `ListCustomers`․
//
// Access the customer list with `ListCustomersResponse.Customers`․
type ListCustomersResponse struct {
	Customers []string `json:"customers"`
}

// ListCustomers lists existing customer IDs in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#get-list-customers for more details․
func (client *Config) ListCustomers() (*ListCustomersResponse, error) {
	res, err := client.get("customers", nil)
	if err != nil {
		return nil, err
	}

	customers := new(ListCustomersResponse)
	if err := json.Unmarshal(res.([]byte), &customers); err != nil {
		return nil, err
	}

	return customers, nil
}

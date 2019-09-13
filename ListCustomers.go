package taxjar

import "encoding/json"

// ListCustomersResponse - TODO (document this)
type ListCustomersResponse struct {
	Customers []string `json:"customers"`
}

// ListCustomers - TODO (document this)
func (client *Config) ListCustomers() (*ListCustomersResponse, error) {
	res, err := client.get("customers")
	if err != nil {
		return nil, err
	}
	customers := new(ListCustomersResponse)
	json.Unmarshal(res.([]byte), &customers)
	return customers, nil
}

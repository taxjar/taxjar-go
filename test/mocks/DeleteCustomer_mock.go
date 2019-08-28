package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// DeleteCustomer - mock response
var DeleteCustomer = new(taxjar.DeleteCustomerResponse)
var _ = json.Unmarshal([]byte(DeleteCustomerJSON), &DeleteCustomer)

// DeleteCustomerJSON - mock DeleteCustomer JSON
var DeleteCustomerJSON = `{
  "customer": {
    "customer_id": "123",
    "exemption_type": "wholesale",
    "exempt_regions": [
      {
        "country": "US",
        "state": "PA"
      }
    ],
    "name": "Test Customer 1",
    "country": "US",
    "state": "TX",
    "zip": "78744",
    "city": "Austin",
    "street": "4120 Freidrich Lane"
  }
}`

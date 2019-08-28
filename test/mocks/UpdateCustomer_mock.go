package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// UpdateCustomer - mock response
var UpdateCustomer = new(taxjar.UpdateCustomerResponse)
var _ = json.Unmarshal([]byte(UpdateCustomerJSON), &UpdateCustomer)

// UpdateCustomerJSON - mock UpdateCustomer JSON
var UpdateCustomerJSON = `{
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

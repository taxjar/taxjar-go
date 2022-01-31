package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// CreateCustomer - mock response
var CreateCustomer = new(taxjar.CreateCustomerResponse)
var _ = json.Unmarshal([]byte(CreateCustomerJSON), &CreateCustomer)

// CreateCustomerJSON - mock CreateCustomer JSON
var CreateCustomerJSON = `{
  "customer": {
    "customer_id": "123",
    "exemption_type": "wholesale",
    "exempt_regions": [
      {
        "country": "US",
        "state": "TX"
      }
    ],
    "name": "Initech",
    "country": "US",
    "state": "TX",
    "zip": "78744",
    "city": "Austin",
    "street": "4120 Freidrich Lane"
  }
}`

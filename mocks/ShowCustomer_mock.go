package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// ShowCustomer - mock response
var ShowCustomer = new(taxjar.ShowCustomerResponse)
var _ = json.Unmarshal([]byte(ShowCustomerJSON), &ShowCustomer)

// ShowCustomerJSON - mock ShowCustomer JSON
var ShowCustomerJSON = `{
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

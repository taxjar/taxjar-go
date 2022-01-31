package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// ValidateAddress - mock response
var ValidateAddress = new(taxjar.ValidateAddressResponse)
var _ = json.Unmarshal([]byte(ValidateAddressJSON), &ValidateAddress)

// ValidateAddressJSON - mock ValidateAddress JSON
var ValidateAddressJSON = `{
  "addresses": [
    {
      "zip": "85297-2176",
      "street": "3301 S Greenfield Rd",
      "state": "AZ",
      "country": "US",
      "city": "Gilbert"
    }
  ]
}`

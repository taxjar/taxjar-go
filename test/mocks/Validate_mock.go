package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// Validate - mock response
var Validate = new(taxjar.ValidateResponse)
var _ = json.Unmarshal([]byte(ValidateJSON), &Validate)

// ValidateJSON - mock Validate JSON
var ValidateJSON = `{
  "validation": {
    "valid": true,
    "exists": true,
    "vies_available": true,
    "vies_response": {
      "country_code": "FR",
      "vat_number": "40303265045",
      "request_date": "2016-02-10",
      "valid": true,
      "name": "SA SODIMAS",
      "address": "11 RUE AMPERE\n26600 PONT DE L ISERE"
    }
  }
}`

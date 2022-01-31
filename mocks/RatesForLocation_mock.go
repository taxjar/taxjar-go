package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// RatesForLocation - mock response
var RatesForLocation = new(taxjar.RatesForLocationResponse)
var _ = json.Unmarshal([]byte(RatesForLocationJSON), &RatesForLocation)

// RatesForLocationJSON - mock RatesForLocation JSON
var RatesForLocationJSON = `{
  "rate": {
    "zip": "89001",
    "state_rate": "0.0",
    "state": "NV",
    "freight_taxable": false,
    "county_rate": "0.071",
    "county": "LINCOLN",
    "country_rate": "0.0",
    "country": "US",
    "combined_rate": "0.071",
    "combined_district_rate": "0.0",
    "city_rate": "0.0",
    "city": "ALAMO"
  }
}`

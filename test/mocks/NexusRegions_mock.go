package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// NexusRegions - mock response
var NexusRegions = new(taxjar.NexusRegionsResponse)
var _ = json.Unmarshal([]byte(NexusRegionsJSON), &NexusRegions)

// NexusRegionsJSON - mock NexusRegions JSON
var NexusRegionsJSON = `{
  "regions": [
    {
      "country_code": "US",
      "country": "United States",
      "region_code": "CA",
      "region": "California"
    },
    {
      "country_code": "US",
      "country": "United States",
      "region_code": "NY",
      "region": "New York"
    },
    {
      "country_code": "US",
      "country": "United States",
      "region_code": "WA",
      "region": "Washington"
    }
  ]
}`

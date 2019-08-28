package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// SummaryRates - mock response
var SummaryRates = new(taxjar.SummaryRatesResponse)
var _ = json.Unmarshal([]byte(SummaryRatesJSON), &SummaryRates)

// SummaryRatesJSON - mock SummaryRates JSON
var SummaryRatesJSON = `{
    "summary_rates": [
        {
            "country_code": "AT",
            "country": "Austria",
            "region_code": null,
            "region": null,
            "minimum_rate": {
                "label": "VAT",
                "rate": 0.2
            },
            "average_rate": {
                "label": "VAT",
                "rate": 0.2
            }
        },
        {
            "country_code": "BE",
            "country": "Belgium",
            "region_code": null,
            "region": null,
            "minimum_rate": {
                "label": "VAT",
                "rate": 0.21
            },
            "average_rate": {
                "label": "VAT",
                "rate": 0.21
            }
        },
        {
            "country_code": "BG",
            "country": "Bulgaria",
            "region_code": null,
            "region": null,
            "minimum_rate": {
                "label": "VAT",
                "rate": 0.2
            },
            "average_rate": {
                "label": "VAT",
                "rate": 0.2
            }
        },
        {
            "country_code": "CA",
            "country": "Canada",
            "region_code": "AB",
            "region": "Alberta",
            "minimum_rate": {
                "label": "GST",
                "rate": 0.05
            },
            "average_rate": {
                "label": "GST",
                "rate": 0.05
            }
        },
        {
            "country_code": "CA",
            "country": "Canada",
            "region_code": "BC",
            "region": "British Columbia",
            "minimum_rate": {
                "label": "GST",
                "rate": 0.05
            },
            "average_rate": {
                "label": "GST/PST",
                "rate": 0.12
            }
        }
    ]
}`

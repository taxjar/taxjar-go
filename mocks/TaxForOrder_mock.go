package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// USTaxForOrder - mock response
var USTaxForOrder = new(taxjar.TaxForOrderResponse)
var _ = json.Unmarshal([]byte(USTaxForOrderJSON), &USTaxForOrder)

// USTaxForOrderJSON - mock TaxForOrder JSON
var USTaxForOrderJSON = `{
  "tax": {
    "exemption_type": "non_exempt",
    "taxable_amount": 15.0,
    "tax_source": "destination",
    "shipping": 1.5,
    "rate": 0.095,
    "order_total_amount": 16.5,
    "jurisdictions": {
      "state": "CA",
      "county": "LOS ANGELES",
      "country": "US",
      "city": "LOS ANGELES"
    },
    "has_nexus": true,
    "freight_taxable": false,
    "breakdown": {
      "taxable_amount": 15.0,
      "tax_collectable": 1.43,
      "state_taxable_amount": 15.0,
      "state_tax_rate": 0.0625,
      "state_tax_collectable": 0.94,
      "special_tax_rate": 0.0225,
      "special_district_taxable_amount": 15.0,
      "special_district_tax_collectable": 0.34,
      "line_items": [
        {
          "taxable_amount": 15.0,
          "tax_collectable": 1.43,
          "state_taxable_amount": 15.0,
          "state_sales_tax_rate": 0.0625,
          "state_amount": 0.94,
          "special_tax_rate": 0.0225,
          "special_district_taxable_amount": 15.0,
          "special_district_amount": 0.34,
          "id": "1",
          "county_taxable_amount": 15.0,
          "county_tax_rate": 0.01,
          "county_amount": 0.15,
          "combined_tax_rate": 0.095,
          "city_taxable_amount": 0.0,
          "city_tax_rate": 0.0,
          "city_amount": 0.0
        }
      ],
      "county_taxable_amount": 15.0,
      "county_tax_rate": 0.01,
      "county_tax_collectable": 0.15,
      "combined_tax_rate": 0.095,
      "city_taxable_amount": 0.0,
      "city_tax_rate": 0.0,
      "city_tax_collectable": 0.0
    },
    "amount_to_collect": 1.43
  }
}`

// CATaxForOrder - mock response
var CATaxForOrder = new(taxjar.TaxForOrderResponse)
var _ = json.Unmarshal([]byte(CATaxForOrderJSON), &CATaxForOrder)

// CATaxForOrderJSON - mock TaxForOrder JSON
var CATaxForOrderJSON = `{
  "tax": {
    "exemption_type": "non_exempt",
    "taxable_amount": 16.5,
    "tax_source": "destination",
    "shipping": 1.5,
    "rate": 0.12,
    "order_total_amount": 16.5,
    "jurisdictions": {
      "state": "BC",
      "country": "CA",
      "city": "VANCOUVER"
    },
    "has_nexus": true,
    "freight_taxable": true,
    "breakdown": {
      "taxable_amount": 16.5,
      "tax_collectable": 1.98,
      "shipping": {
        "taxable_amount": 1.5,
        "tax_collectable": 0.18,
        "qst_taxable_amount": 0.0,
        "qst_tax_rate": 0.0,
        "qst": 0.0,
        "pst_taxable_amount": 1.5,
        "pst_tax_rate": 0.07,
        "pst": 0.11,
        "gst_taxable_amount": 1.5,
        "gst_tax_rate": 0.05,
        "gst": 0.08,
        "combined_tax_rate": 0.12
      },
      "qst_taxable_amount": 0.0,
      "qst_tax_rate": 0.0,
      "qst": 0.0,
      "pst_taxable_amount": 16.5,
      "pst_tax_rate": 0.07,
      "pst": 1.16,
      "line_items": [
        {
          "taxable_amount": 15.0,
          "tax_collectable": 1.8,
          "qst_taxable_amount": 0.0,
          "qst_tax_rate": 0.0,
          "qst": 0.0,
          "pst_taxable_amount": 15.0,
          "pst_tax_rate": 0.07,
          "pst": 1.05,
          "id": "1",
          "gst_taxable_amount": 15.0,
          "gst_tax_rate": 0.05,
          "gst": 0.75,
          "combined_tax_rate": 0.12
        }
      ],
      "gst_taxable_amount": 16.5,
      "gst_tax_rate": 0.05,
      "gst": 0.83,
      "combined_tax_rate": 0.12
    },
    "amount_to_collect": 1.98
  }
}`

// EUTaxForOrder - mock response
var EUTaxForOrder = new(taxjar.TaxForOrderResponse)
var _ = json.Unmarshal([]byte(EUTaxForOrderJSON), &EUTaxForOrder)

// EUTaxForOrderJSON - mock TaxForOrder JSON
var EUTaxForOrderJSON = `{
  "tax": {
    "exemption_type": "non_exempt",
    "taxable_amount": 16.5,
    "tax_source": "destination",
    "shipping": 1.5,
    "rate": 0.2,
    "order_total_amount": 16.5,
    "jurisdictions": {
      "country": "FR",
      "city": "LYON"
    },
    "has_nexus": true,
    "freight_taxable": true,
    "breakdown": {
      "taxable_amount": 16.5,
      "tax_collectable": 3.3,
      "shipping": {
        "taxable_amount": 1.5,
        "tax_collectable": 0.3,
        "country_taxable_amount": 1.5,
        "country_tax_rate": 0.2,
        "country_tax_collectable": 0.3,
        "combined_tax_rate": 0.2
      },
      "line_items": [
        {
          "taxable_amount": 15.0,
          "tax_collectable": 3.0,
          "id": "1",
          "country_taxable_amount": 15.0,
          "country_tax_rate": 0.2,
          "country_tax_collectable": 3.0,
          "combined_tax_rate": 0.2
        }
      ],
      "country_taxable_amount": 16.5,
      "country_tax_rate": 0.2,
      "country_tax_collectable": 3.3,
      "combined_tax_rate": 0.2
    },
    "amount_to_collect": 3.3
  }
}`

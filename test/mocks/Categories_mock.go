package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// Categories - mock response
var Categories = taxjar.CategoriesResponse{
	Categories: []taxjar.Category{
		{
			Name:           "Digital Goods",
			ProductTaxCode: "31000",
			Description:    "Digital products transferred electronically, meaning obtained by the purchaser by means other than tangible storage media.",
		},
		{
			Name:           "Clothing",
			ProductTaxCode: "20010",
			Description:    " All human wearing apparel suitable for general use",
		},
		{
			Name:           "Non-Prescription",
			ProductTaxCode: "51010",
			Description:    "Drugs for human use without a prescription",
		},
	},
}

// CategoriesJSON - mock Categories JSON
var CategoriesJSON, _ = json.Marshal(Categories)

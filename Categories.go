package taxjar

import "encoding/json"

// Category is the structure for a product tax category returned from `Categories` within `CategoriesResponse.Categories`․
type Category struct {
	Name           string `json:"name"`
	ProductTaxCode string `json:"product_tax_code"`
	Description    string `json:"description"`
}

// CategoriesResponse is the structure returned from `Categories`․
//
// Access TaxJar's product tax categories with `CategoriesResponse.Categories`․
type CategoriesResponse struct {
	Categories []Category `json:"categories"`
}

// Categories lists all TaxJar product tax categories to be used for products that are either exempt from sales tax in some jurisdictions or are taxed at reduced rates․
//
// See https://developers.taxjar.com/api/reference/?go#categories for more details․
func (client *Config) Categories() (*CategoriesResponse, error) {
	res, err := client.get("categories", nil)
	if err != nil {
		return nil, err
	}
	categories := new(CategoriesResponse)
	if err := json.Unmarshal(res.([]byte), &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

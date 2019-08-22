package taxjar

import "encoding/json"

// Category - TODO (document this)
type Category struct {
	Name           string `json:"name"`
	ProductTaxCode string `json:"product_tax_code"`
	Description    string `json:"description"`
}

// CategoriesResponse - TODO (document this)
type CategoriesResponse struct {
	Categories []Category `json:"categories"`
}

// Categories - TODO (document this)
func (client *Config) Categories() (*CategoriesResponse, error) {
	res, err := client.get("categories")
	if err != nil {
		return nil, err
	}
	categories := new(CategoriesResponse)
	json.Unmarshal(res.([]byte), &categories)
	return categories, nil
}

package taxjar

import "encoding/json"

// ValidateParams - TODO (document this)
type ValidateParams struct {
	VAT string `url:"vat,omitempty"`
}

// VIESResponse - TODO (document this)
type VIESResponse struct {
	CountryCode string `json:"country_code"`
	VATNumber   string `json:"vat_number"`
	RequestDate string `json:"request_date"`
	Valid       bool   `json:"valid"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

// Validation - TODO (document this)
type Validation struct {
	Valid         bool         `json:"valid"`
	Exists        bool         `json:"exists"`
	VIESAvailable bool         `json:"vies_available"`
	VIESResponse  VIESResponse `json:"vies_response"`
}

// ValidateResponse - TODO (document this)
type ValidateResponse struct {
	Validation Validation `json:"validation"`
}

// Validate - TODO (document this)
func (client *Config) Validate(params ValidateParams) (*ValidateResponse, error) {
	res, err := client.get("validation", params)
	if err != nil {
		return nil, err
	}
	validation := new(ValidateResponse)
	json.Unmarshal(res.([]byte), &validation)
	return validation, nil
}

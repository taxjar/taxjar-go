package taxjar

import "encoding/json"

// ValidateParams should be passed to `Validate` to validate a VAT identification number with VIES (http://ec.europa.eu/taxation_customs/vies/)․
type ValidateParams struct {
	VAT string `url:"vat,omitempty"`
}

// VIESResponse is the structure for a response from VIES (http://ec.europa.eu/taxation_customs/vies/) returned as `ValidateResponse.Validation.ViesResponse`․
type VIESResponse struct {
	CountryCode string `json:"country_code"`
	VATNumber   string `json:"vat_number"`
	RequestDate string `json:"request_date"`
	Valid       bool   `json:"valid"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

// Validation is the structure for a VAT identification number validation returned within `ValidateResponse`․
type Validation struct {
	Valid         bool         `json:"valid"`
	Exists        bool         `json:"exists"`
	VIESAvailable bool         `json:"vies_available"`
	VIESResponse  VIESResponse `json:"vies_response"`
}

// ValidateResponse is the structure returned from `Validate`․
//
// Access the validation with `ValidateResponse.Validation`․
type ValidateResponse struct {
	Validation Validation `json:"validation"`
}

// Validate validates a VAT identification number with VIES (http://ec.europa.eu/taxation_customs/vies/)․
//
// See https://developers.taxjar.com/api/reference/?go#get-validate-a-vat-number for more details․
func (client *Config) Validate(params ValidateParams) (*ValidateResponse, error) {
	res, err := client.get("validation", params)
	if err != nil {
		return nil, err
	}
	validation := new(ValidateResponse)
	if err := json.Unmarshal(res.([]byte), &validation); err != nil {
		return nil, err
	}
	return validation, nil
}

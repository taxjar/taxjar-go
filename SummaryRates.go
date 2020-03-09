package taxjar

import "encoding/json"

// SummaryRate is the structure for a location's summarized (minimum and average) backup rates returned within `SummaryRatesResponse.SummaryRates`․
type SummaryRate struct {
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	RegionCode  string `json:"region_code"`
	Region      string `json:"region"`
	MinimumRate struct {
		Label string  `json:"label"`
		Rate  float64 `json:"rate"`
	} `json:"minimum_rate"`
	AverageRate struct {
		Label string  `json:"label"`
		Rate  float64 `json:"rate"`
	} `json:"average_rate"`
}

// SummaryRatesResponse is the structure returned from `SummaryRates`․
//
// Access the summarized (minimum and average) backup rates with `SummaryRatesResponse.SummaryRates`.
type SummaryRatesResponse struct {
	SummaryRates []SummaryRate `json:"summary_rates"`
}

// SummaryRates retrieves minimum and average sales tax rates by region, which you can use as a backup․
//
// See https://developers.taxjar.com/api/reference/?go#get-summarize-tax-rates-for-all-regions for more details․
func (client *Config) SummaryRates() (*SummaryRatesResponse, error) {
	res, err := client.get("summary_rates", nil)
	if err != nil {
		return nil, err
	}
	summaryRates := new(SummaryRatesResponse)
	if err := json.Unmarshal(res.([]byte), &summaryRates); err != nil {
		return nil, err
	}
	return summaryRates, nil
}

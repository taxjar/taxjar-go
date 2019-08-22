package taxjar

import "encoding/json"

// SummaryRate - TODO (document this)
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

// SummaryRatesResponse - TODO (document this)
type SummaryRatesResponse struct {
	SummaryRates []SummaryRate `json:"summary_rates"`
}

// SummaryRates - TODO (document this)
func (client *Config) SummaryRates() (*SummaryRatesResponse, error) {
	res, err := client.get("summary_rates")
	if err != nil {
		return nil, err
	}
	summaryRates := new(SummaryRatesResponse)
	json.Unmarshal(res.([]byte), &summaryRates)
	return summaryRates, nil
}

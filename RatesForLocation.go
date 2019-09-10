package taxjar

import "encoding/json"

// RatesForLocationParams - TODO (document this)
type RatesForLocationParams struct {
	Country string `url:"country,omitempty"`
	State   string `url:"state,omitempty"`
	City    string `url:"city,omitempty"`
	Street  string `url:"street,omitempty"`
}

// Rate - TODO (document this)
type Rate struct {
	Zip                   string `json:"zip"`
	Country               string `json:"country"`
	Name                  string `json:"name"`
	StandardRate          string `json:"standard_rate"`
	ReducedRate           string `json:"reduced_rate"`
	SuperReducedRate      string `json:"super_reduced_rate"`
	ParkingRate           string `json:"parking_rate"`
	DistanceSaleThreshold string `json:"distance_sale_threshold"`
	CountryRate           string `json:"country_rate"`
	State                 string `json:"state"`
	StateRate             string `json:"state_rate"`
	County                string `json:"county"`
	CountyRate            string `json:"county_rate"`
	City                  string `json:"city"`
	CityRate              string `json:"city_rate"`
	CombinedDistrictRate  string `json:"combined_district_rate"`
	CombinedRate          string `json:"combined_rate"`
	FreightTaxable        bool   `json:"freight_taxable"`
}

// RatesForLocationResponse - TODO (document this)
type RatesForLocationResponse struct {
	Rate Rate `json:"rate"`
}

// RatesForLocation - TODO (document this)
func (client *Config) RatesForLocation(zip string, params ...RatesForLocationParams) (*RatesForLocationResponse, error) {
	res, err := client.get("rates/"+zip, params)
	if err != nil {
		return nil, err
	}
	rate := new(RatesForLocationResponse)
	json.Unmarshal(res.([]byte), &rate)
	return rate, nil
}

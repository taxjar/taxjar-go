package taxjar

import "encoding/json"

// RatesForLocationParams - TODO (document this)
type RatesForLocationParams struct {
	Country string `url:"country,omitempty"`
	State   string `url:"state,omitempty"`
	City    string `url:"city,omitempty"`
	Street  string `url:"street,omitempty"`
}

// RatesForLocationResponse - TODO (document this)
type RatesForLocationResponse struct {
	Rate struct {
		Zip                   string  `json:"zip"`
		Country               string  `json:"country"`
		Name                  string  `json:"name"`
		StandardRate          float64 `json:"standard_rate"`
		ReducedRate           float64 `json:"reduced_rate"`
		SuperReducedRate      float64 `json:"super_reduced_rate"`
		ParkingRate           float64 `json:"parking_rate"`
		DistanceSaleThreshold float64 `json:"distance_sale_threshold"`
		CountryRate           float64 `json:"country_rate"`
		State                 string  `json:"state"`
		StateRate             float64 `json:"state_rate"`
		County                string  `json:"county"`
		CountyRate            float64 `json:"county_rate"`
		City                  string  `json:"city"`
		CityRate              float64 `json:"city_rate"`
		CombinedDistrictRate  float64 `json:"combined_district_rate"`
		CombinedRate          float64 `json:"combined_rate"`
		FreightTaxable        bool    `json:"freight_taxable"`
	} `json:"rate"`
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

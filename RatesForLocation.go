package taxjar

import "encoding/json"

// RatesForLocationParams should be passed to `RatesForLocation` to show the sales tax rates for a given location․
type RatesForLocationParams struct {
	Country string `url:"country,omitempty"`
	State   string `url:"state,omitempty"`
	City    string `url:"city,omitempty"`
	Street  string `url:"street,omitempty"`
}

// Rate is the structure for a given location's sales tax rates․
type Rate struct {
	Zip                   string  `json:"zip"`
	Country               string  `json:"country"`
	Name                  string  `json:"name"`
	StandardRate          float64 `json:"standard_rate,string"`
	ReducedRate           float64 `json:"reduced_rate,string"`
	SuperReducedRate      float64 `json:"super_reduced_rate,string"`
	ParkingRate           float64 `json:"parking_rate,string"`
	DistanceSaleThreshold float64 `json:"distance_sale_threshold,string"`
	CountryRate           float64 `json:"country_rate,string"`
	State                 string  `json:"state"`
	StateRate             float64 `json:"state_rate,string"`
	County                string  `json:"county"`
	CountyRate            float64 `json:"county_rate,string"`
	City                  string  `json:"city"`
	CityRate              float64 `json:"city_rate,string"`
	CombinedDistrictRate  float64 `json:"combined_district_rate,string"`
	CombinedRate          float64 `json:"combined_rate,string"`
	FreightTaxable        bool    `json:"freight_taxable"`
}

// RatesForLocationResponse is the structure returned from `RatesForLocation`․
//
// Access the location's rates with `RatesForLocationResponse.Rate`․
type RatesForLocationResponse struct {
	Rate Rate `json:"rate"`
}

// RatesForLocation shows the sales tax rates for a given location․
//
// Please note `RatesForLocation` only returns the full combined rate for a given location․ It does not support nexus determination, sourcing based on a ship from and ship to address, shipping taxability, product exemptions, customer exemptions, or sales tax holidays․
//
// We recommend using `TaxForOrder` to accurately calculate sales tax for an order․
//
// See https://developers.taxjar.com/api/reference/?go#get-show-tax-rates-for-a-location for more details.
func (client *Config) RatesForLocation(zip string, params ...RatesForLocationParams) (*RatesForLocationResponse, error) {
	var p interface{}
	if len(params) > 0 {
		p = params[0]
	}
	res, err := client.get("rates/"+zip, p)
	if err != nil {
		return nil, err
	}
	rate := new(RatesForLocationResponse)
	json.Unmarshal(res.([]byte), &rate)
	return rate, nil
}

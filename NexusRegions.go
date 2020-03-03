package taxjar

import "encoding/json"

// NexusRegion is the structure for a nexus region returned within `NexusRegionsResponse.Regions`․
type NexusRegion struct {
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	RegionCode  string `json:"region_code"`
	Region      string `json:"region"`
}

// NexusRegionsResponse is the structure returned from `NexusRegions`․
//
// Access the nexus list with `NexusRegionsResponse.Regions`.
type NexusRegionsResponse struct {
	Regions []NexusRegion `json:"regions"`
}

// NexusRegions lists existing nexus locations for a TaxJar account․
//
// See https://developers.taxjar.com/api/reference/?go#get-list-nexus-regions for more details․
func (client *Config) NexusRegions() (*NexusRegionsResponse, error) {
	res, err := client.get("nexus/regions", nil)
	if err != nil {
		return nil, err
	}
	regions := new(NexusRegionsResponse)
	json.Unmarshal(res.([]byte), &regions)
	return regions, nil
}

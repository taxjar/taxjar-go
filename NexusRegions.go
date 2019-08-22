package taxjar

import "encoding/json"

// NexusRegion - TODO (document this)
type NexusRegion struct {
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	RegionCode  string `json:"region_code"`
	Region      string `json:"region"`
}

// NexusRegionsResponse - TODO (document this)
type NexusRegionsResponse struct {
	Regions []NexusRegion `json:"regions"`
}

// NexusRegions - TODO (document this)
func (client *Config) NexusRegions() (*NexusRegionsResponse, error) {
	res, err := client.get("nexus/regions")
	if err != nil {
		return nil, err
	}
	regions := new(NexusRegionsResponse)
	json.Unmarshal(res.([]byte), &regions)
	return regions, nil
}

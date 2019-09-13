package taxjar

// NexusAddress is the structure for a nexus address passed within `TaxForOrderParams.NexusAddresses`․
type NexusAddress struct {
	ID      string `json:"id,omitempty"`
	Country string `json:"country,omitempty"`
	Zip     string `json:"zip,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
}

// TaxLineItem is the structure for a line item passed within `TaxForOrderParams.LineItems`․
type TaxLineItem struct {
	ID             string  `json:"id,omitempty"`
	Quantity       int     `json:"quantity,omitempty"`
	ProductTaxCode string  `json:"product_tax_code,omitempty"`
	UnitPrice      float64 `json:"unit_price,omitempty"`
	Discount       float64 `json:"discount,omitempty"`
}

// TaxForOrderParams should be passed to `TaxForOrder` to calculate tax․
type TaxForOrderParams struct {
	FromCountry    string         `json:"from_country,omitempty"`
	FromZip        string         `json:"from_zip,omitempty"`
	FromState      string         `json:"from_state,omitempty"`
	FromCity       string         `json:"from_city,omitempty"`
	FromStreet     string         `json:"from_street,omitempty"`
	ToCountry      string         `json:"to_country,omitempty"`
	ToZip          string         `json:"to_zip,omitempty"`
	ToState        string         `json:"to_state,omitempty"`
	ToCity         string         `json:"to_city,omitempty"`
	ToStreet       string         `json:"to_street,omitempty"`
	Amount         float64        `json:"amount,omitempty"`
	Shipping       float64        `json:"shipping"`
	CustomerID     string         `json:"customer_id,omitempty"`
	ExemptionType  string         `json:"exemption_type,omitempty"`
	NexusAddresses []NexusAddress `json:"nexus_addresses,omitempty"`
	LineItems      []TaxLineItem  `json:"line_items,omitempty"`
}

// Jurisdictions is the structure for `TaxForOrderResponse.Tax.Jurisdictions`․
type Jurisdictions struct {
	Country string `json:"country"`
	State   string `json:"state"`
	County  string `json:"county"`
	City    string `json:"city"`
}

// Shipping is the structure for `TaxForOrderResponse.Tax.Breakdown.Shipping`․
type Shipping struct {
	TaxableAmount         float64 `json:"taxable_amount"`
	TaxCollectable        float64 `json:"tax_collectable"`
	CombinedTaxRate       float64 `json:"combined_tax_rate"`
	StateTaxableAmount    float64 `json:"state_taxable_amount"`
	StateSalesTaxRate     float64 `json:"state_sales_tax_rate"`
	StateAmount           float64 `json:"state_amount"`
	CountyTaxableAmount   float64 `json:"county_taxable_amount"`
	CountyTaxRate         float64 `json:"county_tax_rate"`
	CountyAmount          float64 `json:"county_amount"`
	CityTaxableAmount     float64 `json:"city_taxable_amount"`
	CityTaxRate           float64 `json:"city_tax_rate"`
	CityAmount            float64 `json:"city_amount"`
	SpecialTaxableAmount  float64 `json:"special_taxable_amount"`
	SpecialTaxRate        float64 `json:"special_tax_rate"`
	SpecialDistrictAmount float64 `json:"special_district_amount"`
	// Canada
	GSTTaxableAmount float64 `json:"gst_taxable_amount"`
	GSTTaxRate       float64 `json:"gst_tax_rate"`
	GST              float64 `json:"gst"`
	PSTTaxableAmount float64 `json:"pst_taxable_amount"`
	PSTTaxRate       float64 `json:"pst_tax_rate"`
	PST              float64 `json:"pst"`
	QSTTaxableAmount float64 `json:"qst_taxable_amount"`
	QSTTaxRate       float64 `json:"qst_tax_rate"`
	QST              float64 `json:"qst"`
	// Other International Attributes
	CountryTaxableAmount  float64 `json:"country_taxable_amount"`
	CountryTaxRate        float64 `json:"country_tax_rate"`
	CountryTaxCollectable float64 `json:"country_tax_collectable"`
}

// LineItemBreakdown is the structure for a line item in `TaxForOrderResponse.Tax.Breakdown.LineItems`․
type LineItemBreakdown struct {
	ID                           string  `json:"id"`
	TaxableAmount                float64 `json:"taxable_amount"`
	TaxCollectable               float64 `json:"tax_collectable"`
	CombinedTaxRate              float64 `json:"combined_tax_rate"`
	StateTaxableAmount           float64 `json:"state_taxable_amount"`
	StateSalesTaxRate            float64 `json:"state_sales_tax_rate"`
	StateAmount                  float64 `json:"state_amount"`
	CountyTaxableAmount          float64 `json:"county_taxable_amount"`
	CountyTaxRate                float64 `json:"county_tax_rate"`
	CountyAmount                 float64 `json:"county_amount"`
	CityTaxableAmount            float64 `json:"city_taxable_amount"`
	CityTaxRate                  float64 `json:"city_tax_rate"`
	CityAmount                   float64 `json:"city_amount"`
	SpecialDistrictTaxableAmount float64 `json:"special_district_taxable_amount"`
	SpecialTaxRate               float64 `json:"special_tax_rate"`
	SpecialDistrictAmount        float64 `json:"special_district_amount"`
	// Canada
	GSTTaxableAmount float64 `json:"gst_taxable_amount"`
	GSTTaxRate       float64 `json:"gst_tax_rate"`
	GST              float64 `json:"gst"`
	PSTTaxableAmount float64 `json:"pst_taxable_amount"`
	PSTTaxRate       float64 `json:"pst_tax_rate"`
	PST              float64 `json:"pst"`
	QSTTaxableAmount float64 `json:"qst_taxable_amount"`
	QSTTaxRate       float64 `json:"qst_tax_rate"`
	QST              float64 `json:"qst"`
	// Other International Attributes
	CountryTaxableAmount  float64 `json:"country_taxable_amount"`
	CountryTaxRate        float64 `json:"country_tax_rate"`
	CountryTaxCollectable float64 `json:"country_tax_collectable"`
}

// Breakdown is the structure for `TaxForOrderResponse.Tax.Breakdown`․
type Breakdown struct {
	TaxableAmount                 float64             `json:"taxable_amount"`
	TaxCollectable                float64             `json:"tax_collectable"`
	CombinedTaxRate               float64             `json:"combined_tax_rate"`
	StateTaxableAmount            float64             `json:"state_taxable_amount"`
	StateTaxRate                  float64             `json:"state_tax_rate"`
	StateTaxCollectable           float64             `json:"state_tax_collectable"`
	CountyTaxableAmount           float64             `json:"county_taxable_amount"`
	CountyTaxRate                 float64             `json:"county_tax_rate"`
	CountyTaxCollectable          float64             `json:"county_tax_collectable"`
	CityTaxableAmount             float64             `json:"city_taxable_amount"`
	CityTaxRate                   float64             `json:"city_tax_rate"`
	CityTaxCollectable            float64             `json:"city_tax_collectable"`
	SpecialDistrictTaxableAmount  float64             `json:"special_district_taxable_amount"`
	SpecialTaxRate                float64             `json:"special_tax_rate"`
	SpecialDistrictTaxCollectable float64             `json:"special_district_tax_collectable"`
	Shipping                      Shipping            `json:"shipping"`
	LineItems                     []LineItemBreakdown `json:"line_items"`
	// Canada
	GSTTaxableAmount float64 `json:"gst_taxable_amount"`
	GSTTaxRate       float64 `json:"gst_tax_rate"`
	GST              float64 `json:"gst"`
	PSTTaxableAmount float64 `json:"pst_taxable_amount"`
	PSTTaxRate       float64 `json:"pst_tax_rate"`
	PST              float64 `json:"pst"`
	QSTTaxableAmount float64 `json:"qst_taxable_amount"`
	QSTTaxRate       float64 `json:"qst_tax_rate"`
	QST              float64 `json:"qst"`
	// Other International Attributes
	CountryTaxableAmount  float64 `json:"country_taxable_amount"`
	CountryTaxRate        float64 `json:"country_tax_rate"`
	CountryTaxCollectable float64 `json:"country_tax_collectable"`
}

// Tax is the stucture for a tax calculation returned within `TaxForOrderResponse`․
type Tax struct {
	OrderTotalAmount float64       `json:"order_total_amount"`
	Shipping         float64       `json:"shipping"`
	TaxableAmount    float64       `json:"taxable_amount"`
	AmountToCollect  float64       `json:"amount_to_collect"`
	Rate             float64       `json:"rate"`
	HasNexus         bool          `json:"has_nexus"`
	FreightTaxable   bool          `json:"freight_taxable"`
	TaxSource        string        `json:"tax_source"`
	ExemptionType    string        `json:"exemption_type"`
	Jurisdictions    Jurisdictions `json:"jurisdictions"`
	Breakdown        Breakdown     `json:"breakdown"`
}

// TaxForOrderResponse is the structure returned from `TaxForOrder`․
//
// Access the calculated tax with `TaxForOrderResponse.Tax`․
type TaxForOrderResponse struct {
	Tax Tax `json:"tax"`
}

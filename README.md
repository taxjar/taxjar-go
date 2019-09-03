# TaxJar Sales Tax API for Go [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/taxjar/taxjar-go?style=flat-square)](CHANGELOG.md) [![GoDoc](https://img.shields.io/badge/godoc-docs-blue.svg?style=flat-square)](https://godoc.org/github.com/taxjar/taxjar-go) [![Build Status](http://img.shields.io/travis/taxjar/taxjar-go.svg?style=flat-square)](https://travis-ci.org/taxjar/taxjar-go) [![Known Vulnerabilities](https://snyk.io/test/github/taxjar/taxjar-go/badge.svg?style=flat-square)](https://snyk.io/test/github/taxjar/taxjar-go)

Official Go client for Sales Tax API v2. For the API documentation, please visit [https://developers.taxjar.com/api/reference/](https://developers.taxjar.com/api/reference/?go).

## Requirements

- [go1.11](https://golang.org/doc/devel/release.html#go1.11) or later.

## Installation

```bash
go get -u github.com/taxjar/taxjar-go
```
```go
// Then, import the package:
import "github.com/taxjar/taxjar-go"
```

## Authentication

[Generate a TaxJar API token](https://app.taxjar.com/api_sign_up/). Enter the token when instantiating a new client. You may want to utilize an environment variable such as `TAXJAR_API_KEY` as seen below:

```go
// Instantiate client with your TaxJar API token:
client := taxjar.NewClient(taxjar.Config{
	APIKey: os.Getenv("TAXJAR_API_KEY"),
})
```
```go
// or configure client after instantiating:
client := taxjar.NewClient()
client.APIKey = os.Getenv("TAXJAR_API_KEY")
```

## Usage

### List all tax categories <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-tax-categories), [GoDoc]())_</small>

> The TaxJar API provides product-level tax rules for a subset of product categories. These categories are to be used for products that are either exempt from sales tax in some jurisdictions or are taxed at reduced rates. You need not pass in a product tax code for sales tax calculations on product that is fully taxable. Simply leave that parameter out.

```go
res, _ := client.Categories()
fmt.Println(res.Categories) // CategoriesResponse.Categories
```

### Calculate sales tax for an order <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-calculate-sales-tax-for-an-order), [GoDoc]())_</small>

> Shows the sales tax that should be collected for a given order.

```go
res, _ := client.TaxForOrder(taxjar.TaxForOrderParams{
	FromCountry: "US",
	FromZip:     "07001",
	FromState:   "NJ",
	ToCountry:   "US",
	ToZip:       "07446",
	ToState:     "NJ",
	Amount:      16.50,
	Shipping:    1.5,
	LineItems:   []taxjar.TaxLineItem{
		{
			Quantity:       1,
			UnitPrice:      15.0,
			ProductTaxCode: "31000",
		},
	},
})
fmt.Println(res.Tax) // TaxForOrderResponse.Tax
fmt.Println(res.Tax.AmountToCollect) // TaxForOrderResponse.Tax.AmountToCollect
```

### List order transactions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-order-transactions), [GoDoc]())_</small>

> Lists existing order transactions created through the API.

```go
res, _ := client.ListOrders(ListOrdersParams{
	FromTransactionDate: "2015/05/01",
	ToTransactionDate:   "2015/05/31",
})
fmt.Println(res.Orders) // ListOrdersResponse.Orders
```

### Show order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-an-order-transaction), [GoDoc]())_</small>

> Shows an existing order transaction created through the API.

```go
res, _ := client.ShowOrder("123")
fmt.Println(res.Order) // ShowOrderResponse.Order
```

### Create order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-create-an-order-transaction), [GoDoc]())_</small>

> Creates a new order transaction.

```go
res, _ := client.CreateOrder(taxjar.CreateOrderParams{
	TransactionID:   "123",
	TransactionDate: "2015/05/14",
	ToCountry:       "US",
	ToZip:           "90002",
	ToState:         "CA",
	ToCity:          "Los Angeles",
	ToStreet:        "123 Palm Grove Ln",
	Amount:          17.45,
	Shipping:        1.5,
	SalesTax:        0.95,
	LineItems:       []taxjar.OrderLineItem{
		{
			Quantity:          1,
			ProductIdentifier: "12-34243-9",
			Description:       "Fuzzy Widget",
			UnitPrice:         15.0,
			SalesTax:          0.95,
		},
	},
})
fmt.Println(res.Order) // CreateOrderResponse.Order
```

### Update order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#put-update-an-order-transaction), [GoDoc]())_</small>

> Updates an existing order transaction created through the API.

```go
res, _ := client.UpdateOrder(taxjar.UpdateOrderParams{
	TransactionID: "123",
	Amount:        17.45,
	Shipping:      1.5,
	LineItems:     []taxjar.OrderLineItem{
		{
			Quantity:          1,
			ProductIdentifier: "12-34243-0",
			Description:       "Heavy Widget",
			UnitPrice:         15.0,
			Discount:          0.0,
			SalesTax:          0.95,
		},
	},
})
fmt.Println(res.Order) // UpdateOrderResponse.Order
```

### Delete order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#delete-delete-an-order-transaction), [GoDoc]())_</small>

> Deletes an existing order transaction created through the API.

```go
res, _ := client.DeleteOrder("123")
fmt.Println(res.Order) // DeleteOrderResponse.Order
```

### List refund transactions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-refund-transactions), [GoDoc]())_</small>

> Lists existing refund transactions created through the API.

```go
res, _ := client.ListRefunds(taxjar.ListRefundsParams{
	FromTransactionDate: "2015/05/01",
	ToTransactionDate:   "2015/05/31",
})
fmt.Println(res.Refunds) // ListRefundsResponse.Refunds
```

### Show refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-a-refund-transaction), [GoDoc]())_</small>

> Shows an existing refund transaction created through the API.

```go
res, _ := client.ShowRefund("321")
fmt.Println(res.Refund) // ShowRefundResponse.Refund
```

### Create refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-create-a-refund-transaction), [GoDoc]())_</small>

> Creates a new refund transaction.

```go
res, _ := client.CreateRefund(taxjar.CreateRefundParams{
	TransactionID:          "123",
	TransactionDate:        "2015/05/14",
	TransactionReferenceID: "123",
	ToCountry:              "US",
	ToZip:                  "90002",
	ToState:                "CA",
	ToCity:                 "Los Angeles",
	ToStreet:               "123 Palm Grove Ln",
	Amount:                 17.45,
	Shipping:               1.5,
	SalesTax:               0.95,
	LineItems:              []taxjar.RefundLineItem{
		{
			Quantity:          1,
			ProductIdentifier: "12-34243-9",
			Description:       "Fuzzy Widget",
			UnitPrice:         15.0,
			SalesTax:          0.95,
		},
	},
})
fmt.Println(res.Refund) // CreateRefundResponse.Refund
```

### Update refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#put-update-a-refund-transaction), [GoDoc]())_</small>

> Updates an existing refund transaction created through the API.

```go
res, _ := client.UpdateRefund(taxjar.UpdateRefundParams{
	TransactionID: "123",
	Amount:        17.95,
	Shipping:      2.0,
	LineItems:     []taxjar.RefundLineItem{
		{
			Quantity:          1,
			ProductIdentifier: "12-34243-0",
			Description:       "Heavy Widget",
			UnitPrice:         15.0,
			SalesTax:          0.95,
		},
	},
})
fmt.Println(res.Refund) // UpdateRefundResponse.Refund
```

### Delete refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#delete-delete-a-refund-transaction), [GoDoc]())_</small>

> Deletes an existing refund transaction created through the API.

```go
res, _ := client.DeleteRefund("123")
fmt.Println(res.Refund) // DeleteRefundResponse.Refund
```

### List customers <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-customers), [GoDoc]())_</small>

> Lists existing customers created through the API.

```go
res, _ := client.ListCustomers()
fmt.Println(res.Customers) // ListCustomersResponse.Customers
```

### Show customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-a-customer), [GoDoc]())_</small>

> Shows an existing customer created through the API.

```go
res, _ := client.ShowCustomer("123")
fmt.Println(res.Customer) // ShowCustomerResponse.Customer
```

### Create customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-create-a-customer), [GoDoc]())_</small>

> Creates a new customer.

```go
res, _ := client.CreateCustomer(taxjar.CreateCustomerParams{
	CustomerID:    "123",
	ExemptionType: "wholesale",
	Name:          "Dunder Mifflin Paper Company",
	ExemptRegions: []taxjar.ExemptRegion{
		{
			Country: "US",
			State:   "FL",
		},
		{
			Country: "US",
			State:   "PA",
		},
	},
	Country:       "US",
	State:         "PA",
	Zip:           "18504",
	City:          "Scranton",
	Street:        "1725 Slough Avenue",
})
fmt.Println(res.Customer) // CreateCustomerResponse.Customer
```

### Update customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#put-update-a-customer), [GoDoc]())_</small>

> Updates an existing customer created through the API.

```go
res, _ := client.UpdateCustomer(taxjar.UpdateCustomerParams{
	CustomerID:    "123",
	ExemptionType: "wholesale",
	Name:          "Sterling Cooper",
	ExemptRegions: []taxjar.ExemptRegion{
		{
			Country: "US",
			State:   "NY",
		},
	},
	Country:       "US",
	State:         "NY",
	Zip:           "10010",
	City:          "New York",
	Street:        "405 Madison Ave",
})
fmt.Println(res.Customer) // UpdateCustomerResponse.Customer
```

### Delete customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#delete-delete-a-customer), [GoDoc]())_</small>

> Deletes an existing customer created through the API.

```go
res, _ := client.DeleteCustomer("123")
fmt.Println(res.Customer) // DeleteCustomerResponse.Customer
```

### List tax rates for a location (by zip/postal code) <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-tax-rates-for-a-location), [GoDoc]())_</small>

> Shows the sales tax rates for a given location.
>
> **Please note this method only returns the full combined rate for a given location.** It does not support nexus determination, sourcing based on a ship from and ship to address, shipping taxability, product exemptions, customer exemptions, or sales tax holidays. We recommend using [`TaxForOrder` to accurately calculate sales tax for an order](#calculate-sales-tax-for-an-order-smallAPI-docs-godocsmall)).

```go
res, _ := client.RatesForLocation("90002")
fmt.Println(res.Rate) // RatesForLocationResponse.Rate
```

### List nexus regions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-nexus-regions), [GoDoc]())_</small>

> Lists existing nexus locations for a TaxJar account.

```go
res, _ := client.NexusRegions()
fmt.Println(res.Regions) // NexusRegionsResponse.Regions
```

### Validate an address <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-validate-an-address), [GoDoc]())_</small>

> Validates a customer address and returns back a collection of address matches. **Address validation requires a [TaxJar Plus](https://www.taxjar.com/plus/) subscription.**

```go
res, _ := client.ValidateAddress(taxjar.ValidateAddressParams{
	Country: "US",
	State:   "AZ",
	Zip:     "85297",
	City:    "Gilbert",
	Street:  "3301 South Greenfield Rd",
})
fmt.Println(res.Addresses) // ValidateAddressResponse.Addresses
```

### Validate a VAT number <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-validate-a-vat-number), [GoDoc]())_</small>

> Validates an existing VAT identification number against [VIES](http://ec.europa.eu/taxation_customs/vies/).

```go
res, _ := client.Validate({
	VAT: "FR40303265045",
})
fmt.Println(res.Validation) // ValidateResponse.Validation
```

### Summarize tax rates for all regions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-summarize-tax-rates-for-all-regions), [GoDoc]())_</small>

> Retrieve minimum and average sales tax rates by region as a backup.

```go
res, _ := client.SummaryRates()
fmt.Println(res.SummaryRates) // SummaryRatesResponse.SummaryRates
```

## Sandbox Environment

You can easily configure the client to use the [TaxJar Sandbox](https://developers.taxjar.com/api/reference/?go#sandbox-environment):

```go
import "github.com/taxjar/taxjar-go"

// Instantiate client and set `APIURL`:
client := taxjar.NewClient(taxjar.Config{
	APIKey: os.Getenv("TAXJAR_API_KEY"),
	APIURL: taxjar.SandboxAPIURL,
})

// or
client := taxjar.NewClient()
client.APIKey = os.Getenv("TAXJAR_API_KEY")
client.APIURL = taxjar.SandboxAPIURL
```

For testing specific [error response codes](https://developers.taxjar.com/api/reference/?go#errors), pass the custom `X-TJ-Expected-Response` header:

```javascript
client.Headers = map[string]interface{}{
	"X-TJ-Expected-Response": 422,
}
```

## Error Handling

```go
res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
	FromCountry: "US",
	FromZip:     "07001",
	FromState:   "NJ",
	ToCountry:   "US",
	ToZip:       "07446",
	ToState:     "NJ",
	Amount:      16.50,
	Shipping:    1.5,
	LineItems:   []taxjar.TaxLineItem{
		{
			Quantity:       1,
			UnitPrice:      15.0,
			ProductTaxCode: "31000",
		},
	},
})
if err != nil {
  fmt.Println(err) // taxjar: 401 Unauthorized - Not Authorized for route 'POST /v2/taxes'
} else {
  fmt.Println(res.Tax)
}
// or extract more information by asserting to `*taxjar.Error`
if err := err.(*taxjar.Error); err != nil {
	fmt.Println(err.Status) // 401
	fmt.Println(err.Err) // Unauthorized
	fmt.Println(err.Detail) // Not authorized for route `POST /v2/taxes'
	fmt.Printf("%+v", errors.Wrap(err, "")) // Stack trace:
	// taxjar: 401 Unauthorized - Not Authorized for route 'POST /v2/taxes'
	//
	// main.main
	//         /Path/to/your/file.go:292
	// runtime.main
	//         /usr/local/go/src/runtime/proc.go:200
	// runtime.goexit
	//         /usr/local/go/src/runtime/asm_amd64.s:1337
} else {
	fmt.Println(res.Tax)
}
```

## Testing

```
make test
```

To validate API methods in the TaxJar sandbox environment, pass the following environment variables:

```bash
TAXJAR_API_URL="https://api.sandbox.taxjar.com" \
TAXJAR_API_KEY="9e0cd62a22f451701f29c3bde214" \
make test
```

# TaxJar Sales Tax API for Go ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/taxjar/taxjar-go?style=flat-square&label=release&sort=semver) [![GoDoc](https://img.shields.io/badge/godoc-docs-blue.svg?style=flat-square&color=darkturquoise)](https://godoc.org/github.com/taxjar/taxjar-go) [![Build Status](https://img.shields.io/travis/taxjar/taxjar-go.svg?style=flat-square)](https://travis-ci.org/taxjar/taxjar-go)

<a href="https://developers.taxjar.com"><img src="https://www.taxjar.com/img/TJ_logo_color_office_png.png" alt="TaxJar" width="220"></a>

Official Go client for Sales Tax API v2. For the API documentation, please visit [https://developers.taxjar.com/api/reference/](https://developers.taxjar.com/api/reference/?go).

<hr>

[Requirements](#requirements)<br>
[Installation](#installation)<br>
[Authentication](#authentication)<br>
[Usage](#usage)<br>
[Error Handling](#error-handling)<br>
[Optional Configuration](#optional-configuration)<br>
[Sandbox Environment](#sandbox-environment)<br>
[Testing](#testing)

<hr>

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

[Generate a TaxJar API token](https://app.taxjar.com/api_sign_up/). Enter the token when instantiating with [`NewClient`](https://godoc.org/github.com/taxjar/taxjar-go/#NewClient). You may want to utilize an environment variable such as `TAXJAR_API_KEY` as seen below:

**func [NewClient](https://godoc.org/github.com/taxjar/taxjar-go/#NewClient)(config ...[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)**

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

You're now ready to use TaxJar! [Check out our quickstart guide](https://developers.taxjar.com/api/guides/go/#go-quickstart) to get up and running quickly.

## Usage

[`Categories` - List all tax categories](#list-all-tax-categories-api-docs)<br>
[`TaxForOrder` - Calculate sales tax for an order](#calculate-sales-tax-for-an-order-api-docs)<br>
[`ListOrders` - List order transactions](#list-order-transactions-api-docs)<br>
[`ShowOrder` - Show order transaction](#show-order-transaction-api-docs)<br>
[`CreateOrder` - Create order transaction](#create-order-transaction-api-docs)<br>
[`UpdateOrder` - Update order transaction](#update-order-transaction-api-docs)<br>
[`DeleteOrder` - Delete order transaction](#delete-order-transaction-api-docs)<br>
[`ListRefunds` - List refund transactions](#list-refund-transactions-api-docs)<br>
[`ShowRefund` - Show refund transaction](#show-refund-transaction-api-docs)<br>
[`CreateRefund` - Create refund transaction](#create-refund-transaction-api-docs)<br>
[`UpdateRefund` - Update refund transaction](#update-refund-transaction-api-docs)<br>
[`DeleteRefund` - Delete refund transaction](#delete-refund-transaction-api-docs)<br>
[`ListCustomers` - List customers](#list-customers-api-docs)<br>
[`ShowCustomer` - Show customer](#show-customer-api-docs)<br>
[`CreateCustomer` - Create customer](#create-customer-api-docs)<br>
[`UpdateCustomer` - Update customer](#update-customer-api-docs)<br>
[`DeleteCustomer` - Delete customer](#delete-customer-api-docs)<br>
[`RatesForLocation` - List tax rates for a location (by zip/postal code)](#list-tax-rates-for-a-location-by-zippostal-code-api-docs)<br>
[`NexusRegions` - List nexus regions](#list-nexus-regions-api-docs)<br>
[`ValidateAddress` - Validate an address](#validate-an-address-api-docs)<br>
[`Validate` - Validate a VAT number](#validate-a-vat-number-api-docs)<br>
[`SummaryRates` - Summarize tax rates for all regions](#summarize-tax-rates-for-all-regions-api-docs)

<hr>

### List all tax categories <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-tax-categories))_</small>

> The TaxJar API provides product-level tax rules for a subset of product categories. These categories are to be used for products that are either exempt from sales tax in some jurisdictions or are taxed at reduced rates. You need not pass in a product tax code for sales tax calculations on product that is fully taxable. Simply leave that parameter out.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [Categories](https://godoc.org/github.com/taxjar/taxjar-go/#Config.Categories)() (\*[CategoriesResponse](https://godoc.org/github.com/taxjar/taxjar-go/#CategoriesResponse), [error](https://godoc.org/builtin#error))**

```go
res, _ := client.Categories()
fmt.Println(res.Categories) // CategoriesResponse.Categories
```

### Calculate sales tax for an order <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-calculate-sales-tax-for-an-order))_</small>

> Shows the sales tax that should be collected for a given order.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [TaxForOrder](https://godoc.org/github.com/taxjar/taxjar-go/#Config.TaxForOrder)(params [TaxForOrderParams](https://godoc.org/github.com/taxjar/taxjar-go/#TaxForOrderParams)) (\*[TaxForOrderResponse](https://godoc.org/github.com/taxjar/taxjar-go/#TaxForOrderResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.TaxForOrder(taxjar.TaxForOrderParams{
	FromCountry: "US",
	FromZip:     "94025",
	FromState:   "CA",
	FromCity:    "Menlo Park",
	FromStreet:  "2825 Sand Hill Rd",
	ToCountry:   "US",
	ToZip:       "94303",
	ToState:     "CA",
	ToCity:      "Palo Alto",
	ToStreet:    "5230 Newell Road",
	Amount:      267.9,
	Shipping:    0,
	LineItems:   []taxjar.TaxLineItem{
		{
			ID:             "1",
			Quantity:       1,
			ProductTaxCode: "19005",
			UnitPrice:      535.8,
			Discount:       267.9,
		},
	},
})
fmt.Println(res.Tax) // TaxForOrderResponse.Tax
fmt.Println(res.Tax.AmountToCollect) // TaxForOrderResponse.Tax.AmountToCollect
```

### List order transactions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-order-transactions))_</small>

> Lists existing order transactions created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ListOrders](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ListOrders)(params [ListOrdersParams](https://godoc.org/github.com/taxjar/taxjar-go/#ListOrdersParams)) (\*[ListOrdersResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ListOrdersResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.ListOrders(ListOrdersParams{
	FromTransactionDate: "2015/05/01",
	ToTransactionDate:   "2015/05/31",
})
fmt.Println(res.Orders) // ListOrdersResponse.Orders
```

### Show order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-an-order-transaction))_</small>

> Shows an existing order transaction created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ShowOrder](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ShowOrder)(transactionID [string](https://godoc.org/builtin/#string), params ...[ShowOrderParams](https://godoc.org/github.com/taxjar/taxjar-go/#ShowOrderParams)) (\*[ShowOrderResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ShowOrderResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.ShowOrder("123")
fmt.Println(res.Order) // ShowOrderResponse.Order
```

### Create order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-create-an-order-transaction))_</small>

> Creates a new order transaction.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [CreateOrder](https://godoc.org/github.com/taxjar/taxjar-go/#Config.CreateOrder)(params [CreateOrderParams](https://godoc.org/github.com/taxjar/taxjar-go/#CreateOrderParams)) (\*[CreateOrderResponse](https://godoc.org/github.com/taxjar/taxjar-go/#CreateOrderResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.CreateOrder(taxjar.CreateOrderParams{
	TransactionID:   "123",
	TransactionDate: "2019/05/15",
	FromCountry:     "US",
	FromZip:         "94025",
	FromState:       "CA",
	FromCity:        "Menlo Park",
	FromStreet:      "2825 Sand Hill Rd",
	ToCountry:       "US",
	ToZip:           "94303",
	ToState:         "CA",
	ToCity:          "Palo Alto",
	ToStreet:        "5230 Newell Road",
	Amount:          267.9,
	Shipping:        0,
	SalesTax:        0,
	LineItems:       []taxjar.OrderLineItem{
		{
			ID:             "1",
			Quantity:       1,
			Description:    "Legal Services",
			ProductTaxCode: "19005",
			UnitPrice:      535.8,
			Discount:       267.9,
			SalesTax:       0,
		},
	},
})
fmt.Println(res.Order) // CreateOrderResponse.Order
```

### Update order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#put-update-an-order-transaction))_</small>

> Updates an existing order transaction created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [UpdateOrder](https://godoc.org/github.com/taxjar/taxjar-go/#Config.UpdateOrder)(params [UpdateOrderParams](https://godoc.org/github.com/taxjar/taxjar-go/#UpdateOrderParams)) (\*[UpdateOrderResponse](https://godoc.org/github.com/taxjar/taxjar-go/#UpdateOrderResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.UpdateOrder(taxjar.UpdateOrderParams{
	TransactionID: "123",
	Amount:        283.6,
	Shipping:      5,
	SalesTax:      1.04,
	LineItems:     []taxjar.OrderLineItem{
		{
			ID:             "1",
			Quantity:       1,
			Description:    "Legal Services",
			ProductTaxCode: "19005",
			UnitPrice:      535.8,
			Discount:       267.9,
			SalesTax:       0,
		},
		{
			ID:          "2",
			Quantity:    2,
			Description: "Hoberman Switch Pitch",
			UnitPrice:   10.7,
			Discount:    10.7,
			SalesTax:    1.04,
		},
	},
})
fmt.Println(res.Order) // UpdateOrderResponse.Order
```

### Delete order transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#delete-delete-an-order-transaction))_</small>

> Deletes an existing order transaction created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [DeleteOrder](https://godoc.org/github.com/taxjar/taxjar-go/#Config.DeleteOrder)(transactionID [string](https://godoc.org/builtin/#string), params ...[DeleteOrderParams](https://godoc.org/github.com/taxjar/taxjar-go/#DeleteOrderParams)) (\*[DeleteOrderResponse](https://godoc.org/github.com/taxjar/taxjar-go/#DeleteOrderResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.DeleteOrder("123")
fmt.Println(res.Order) // DeleteOrderResponse.Order
```

### List refund transactions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-refund-transactions))_</small>

> Lists existing refund transactions created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ListRefunds](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ListRefunds)(params [ListRefundsParams](https://godoc.org/github.com/taxjar/taxjar-go/#ListRefundsParams)) (\*[ListRefundsResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ListRefundsResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.ListRefunds(taxjar.ListRefundsParams{
	FromTransactionDate: "2015/05/01",
	ToTransactionDate:   "2015/05/31",
})
fmt.Println(res.Refunds) // ListRefundsResponse.Refunds
```

### Show refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-a-refund-transaction))_</small>

> Shows an existing refund transaction created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ShowRefund](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ShowRefund)(transactionID [string](https://godoc.org/builtin/#string), params ...[ShowRefundParams](https://godoc.org/github.com/taxjar/taxjar-go/#ShowRefundParams)) (\*[ShowRefundResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ShowRefundResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.ShowRefund("321")
fmt.Println(res.Refund) // ShowRefundResponse.Refund
```

### Create refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-create-a-refund-transaction))_</small>

> Creates a new refund transaction.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [CreateRefund](https://godoc.org/github.com/taxjar/taxjar-go/#Config.CreateRefund)(params [CreateRefundParams](https://godoc.org/github.com/taxjar/taxjar-go/#CreateRefundParams)) (\*[CreateRefundResponse](https://godoc.org/github.com/taxjar/taxjar-go/#CreateRefundResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.CreateRefund(taxjar.CreateRefundParams{
	TransactionID:          "123-refund",
	TransactionReferenceID: "123",
	TransactionDate:        "2019/05/15",
	FromCountry:            "US",
	FromZip:                "94025",
	FromState:              "CA",
	FromCity:               "Menlo Park",
	FromStreet:             "2825 Sand Hill Rd",
	ToCountry:              "US",
	ToZip:                  "94303",
	ToState:                "CA",
	ToCity:                 "Palo Alto",
	ToStreet:               "5230 Newell Road",
	Amount:                 -5.35,
	Shipping:               -0,
	SalesTax:               -0.52,
	LineItems:              []taxjar.RefundLineItem{
		{
			ID:             "1",
			Quantity:       1,
			Description:    "Legal Services",
			ProductTaxCode: "19005",
			UnitPrice:      -0,
			Discount:       -0,
			SalesTax:       -0,
		},
		{
			ID:          "2",
			Quantity:    1,
			Description: "Hoberman Switch Pitch",
			UnitPrice:   -0,
			Discount:    -5.35,
			SalesTax:    -0.52,
		},
	},
})
fmt.Println(res.Refund) // CreateRefundResponse.Refund
```

### Update refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#put-update-a-refund-transaction))_</small>

> Updates an existing refund transaction created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [UpdateRefund](https://godoc.org/github.com/taxjar/taxjar-go/#Config.UpdateRefund)(params [UpdateRefundParams](https://godoc.org/github.com/taxjar/taxjar-go/#UpdateRefundParams)) (\*[UpdateRefundResponse](https://godoc.org/github.com/taxjar/taxjar-go/#UpdateRefundResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.UpdateRefund(taxjar.UpdateRefundParams{
	TransactionID:          "123-refund",
	TransactionReferenceID: "123",
	Amount:                 -10.35,
	Shipping:               -5,
})
fmt.Println(res.Refund) // UpdateRefundResponse.Refund
```

### Delete refund transaction <small>_([API docs](https://developers.taxjar.com/api/reference/?go#delete-delete-a-refund-transaction))_</small>

> Deletes an existing refund transaction created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [DeleteRefund](https://godoc.org/github.com/taxjar/taxjar-go/#Config.DeleteRefund)(transactionID [string](https://godoc.org/builtin/#string), params ...[DeleteRefundParams](https://godoc.org/github.com/taxjar/taxjar-go/#DeleteRefundParams)) (\*[DeleteRefundResponse](https://godoc.org/github.com/taxjar/taxjar-go/#DeleteRefundResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.DeleteRefund("123-refund")
fmt.Println(res.Refund) // DeleteRefundResponse.Refund
```

### List customers <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-customers))_</small>

> Lists existing customers created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ListCustomers](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ListCustomers)() (\*[ListCustomersResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ListCustomersResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.ListCustomers()
fmt.Println(res.Customers) // ListCustomersResponse.Customers
```

### Show customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-a-customer))_</small>

> Shows an existing customer created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ShowCustomer](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ShowCustomer)(customerID [string](https://godoc.org/builtin/#string)) (\*[ShowCustomerResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ShowCustomerResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.ShowCustomer("123")
fmt.Println(res.Customer) // ShowCustomerResponse.Customer
```

### Create customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-create-a-customer))_</small>

> Creates a new customer.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [CreateCustomer](https://godoc.org/github.com/taxjar/taxjar-go/#Config.CreateCustomer)(params [CreateCustomerParams](https://godoc.org/github.com/taxjar/taxjar-go/#CreateCustomerParams)) (\*[CreateCustomerResponse](https://godoc.org/github.com/taxjar/taxjar-go/#CreateCustomerResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.CreateCustomer(taxjar.CreateCustomerParams{
	CustomerID:    "123",
	ExemptionType: "wholesale",
	Name:          "Initech",
	ExemptRegions: []taxjar.ExemptRegion{
		{
			Country: "US",
			State:   "TX",
		},
	},
	Country: "US",
	State:   "TX",
	Zip:     "78744",
	City:    "Austin",
	Street:  "4120 Freidrich Lane",
})
fmt.Println(res.Customer) // CreateCustomerResponse.Customer
```

### Update customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#put-update-a-customer))_</small>

> Updates an existing customer created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [UpdateCustomer](https://godoc.org/github.com/taxjar/taxjar-go/#Config.UpdateCustomer)(params [UpdateCustomerParams](https://godoc.org/github.com/taxjar/taxjar-go/#UpdateCustomerParams)) (\*[UpdateCustomerResponse](https://godoc.org/github.com/taxjar/taxjar-go/#UpdateCustomerResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.UpdateCustomer(taxjar.UpdateCustomerParams{
	CustomerID:    "123",
	ExemptionType: "non_exempt",
	Name:          "Initech",
	ExemptRegions: []taxjar.ExemptRegion{
		{
			Country: "US",
			State:   "MA",
		},
		{
			Country: "US",
			State:   "TX",
		},
	},
})
fmt.Println(res.Customer) // UpdateCustomerResponse.Customer
```

### Delete customer <small>_([API docs](https://developers.taxjar.com/api/reference/?go#delete-delete-a-customer))_</small>

> Deletes an existing customer created through the API.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [DeleteCustomer](https://godoc.org/github.com/taxjar/taxjar-go/#Config.DeleteCustomer)(customerID [string](https://godoc.org/builtin/#string)) (\*[DeleteCustomerResponse](https://godoc.org/github.com/taxjar/taxjar-go/#DeleteCustomerResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.DeleteCustomer("123")
fmt.Println(res.Customer) // DeleteCustomerResponse.Customer
```

### List tax rates for a location (by zip/postal code) <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-show-tax-rates-for-a-location))_</small>

> Shows the sales tax rates for a given location.
>
> **Please note this method only returns the full combined rate for a given location.** It does not support nexus determination, sourcing based on a ship from and ship to address, shipping taxability, product exemptions, customer exemptions, or sales tax holidays. We recommend using [`TaxForOrder` to accurately calculate sales tax for an order](#calculate-sales-tax-for-an-order-api-docs).

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [RatesForLocation](https://godoc.org/github.com/taxjar/taxjar-go/#Config.RatesForLocation)(zip [string](https://godoc.org/builtin/#string), params ...[RatesForLocationParams](https://godoc.org/github.com/taxjar/taxjar-go/#RatesForLocationParams)) (\*[RatesForLocationResponse](https://godoc.org/github.com/taxjar/taxjar-go/#RatesForLocationResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.RatesForLocation("90002")
fmt.Println(res.Rate) // RatesForLocationResponse.Rate
```

### List nexus regions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-list-nexus-regions))_</small>

> Lists existing nexus locations for a TaxJar account.

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [NexusRegions](https://godoc.org/github.com/taxjar/taxjar-go/#Config.NexusRegions)() (\*[NexusRegionsResponse](https://godoc.org/github.com/taxjar/taxjar-go/#NexusRegionsResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.NexusRegions()
fmt.Println(res.Regions) // NexusRegionsResponse.Regions
```

### Validate an address <small>_([API docs](https://developers.taxjar.com/api/reference/?go#post-validate-an-address))_</small>

> Validates a customer address and returns back a collection of address matches. **Address validation requires a [TaxJar Plus](https://www.taxjar.com/plus/) subscription.**

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [ValidateAddress](https://godoc.org/github.com/taxjar/taxjar-go/#Config.ValidateAddress)(params [ValidateAddressParams](https://godoc.org/github.com/taxjar/taxjar-go/#ValidateAddressParams)) (\*[ValidateAddressResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ValidateAddressResponse), [error](https://godoc.org/builtin/#error))**

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

### Validate a VAT number <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-validate-a-vat-number))_</small>

> Validates an existing VAT identification number against [VIES](http://ec.europa.eu/taxation_customs/vies/).

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [Validate](https://godoc.org/github.com/taxjar/taxjar-go/#Config.Validate)(params [ValidateParams](https://godoc.org/github.com/taxjar/taxjar-go/#ValidateParams)) (\*[ValidateResponse](https://godoc.org/github.com/taxjar/taxjar-go/#ValidateResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.Validate({
	VAT: "FR40303265045",
})
fmt.Println(res.Validation) // ValidateResponse.Validation
```

### Summarize tax rates for all regions <small>_([API docs](https://developers.taxjar.com/api/reference/?go#get-summarize-tax-rates-for-all-regions))_</small>

> Retrieve minimum and average sales tax rates by region as a backup.
>
> This method is useful for periodically pulling down rates to use if the SmartCalcs API is unavailable. However, it does not support nexus determination, sourcing based on a ship from and ship to address, shipping taxability, product exemptions, customer exemptions, or sales tax holidays. We recommend using [`TaxForOrder` to accurately calculate sales tax for an order](#calculate-sales-tax-for-an-order-api-docs).

**func (client \*[Config](https://godoc.org/github.com/taxjar/taxjar-go/#Config)) [SummaryRates](https://godoc.org/github.com/taxjar/taxjar-go/#Config.SummaryRates)() (\*[SummaryRatesResponse](https://godoc.org/github.com/taxjar/taxjar-go/#SummaryRatesResponse), [error](https://godoc.org/builtin/#error))**

```go
res, _ := client.SummaryRates()
fmt.Println(res.SummaryRates) // SummaryRatesResponse.SummaryRates
```

## Error Handling

```go
res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
	FromCountry: "US",
	FromZip:     "94025",
	FromState:   "CA",
	FromCity:    "Menlo Park",
	FromStreet:  "2825 Sand Hill Rd",
	ToCountry:   "US",
	ToZip:       "94303",
	ToState:     "CA",
	ToCity:      "Palo Alto",
	ToStreet:    "5230 Newell Road",
	Amount:      267.9,
	Shipping:    0,
	LineItems:   []taxjar.TaxLineItem{
		{
			ID:             "1",
			Quantity:       1,
			ProductTaxCode: "19005",
			UnitPrice:      535.8,
			Discount:       267.9,
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

## Optional Configuration

To add additional headers to each request, assign them to `client.Headers`. For example, to test specific [error response codes](https://developers.taxjar.com/api/reference/?go#errors), pass the custom `X-TJ-Expected-Response` header:

```go
client.Headers = map[string]interface{}{
	"X-TJ-Expected-Response": 422,
}
```

If you'd like to customize the timeout for requests, pass a time value to `client.Timeout`.

```go
client.Timeout = 45 * time.Second // taxjar.DefaultTimeout: 30 * time.Second
```

To set more detailed timeouts, you may also pass a custom transport to `client.Transport`.

```go
client.Transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   20 * Time.Second,
		KeepAlive: 20 * Time.Second,
	}).DialContext,
	TLSHandshakeTimeout:   20 * time.Second,
	ExpectContinueTimeout: 8 * time.Second,
	ResponseHeaderTimeout: 6 * time.Second,
}

/* taxjar.DefaultTransport:
&http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   10 * Time.Second,
		KeepAlive: 10 * Time.Second,
	}).DialContext,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 4 * time.Second,
	ResponseHeaderTimeout: 3 * time.Second,
}
*/
```

For even more customization, pass a custom `*http.Client` to `client.HTTPClient`.

```go
client.HTTPClient = &http.Client{/* your configuration here */}
```

## Sandbox Environment

You may also configure the client to use [TaxJar's sandbox environment](https://developers.taxjar.com/api/reference/?go#sandbox-environment). **The sandbox environment requires a [TaxJar Plus](https://www.taxjar.com/plus/) subscription.**

```go
import "github.com/taxjar/taxjar-go"

// Instantiate client and set `APIURL`:
client := taxjar.NewClient(taxjar.Config{
	APIKey: os.Getenv("TAXJAR_SANDBOX_API_KEY"),
	APIURL: taxjar.SandboxAPIURL,
})

// or
client := taxjar.NewClient()
client.APIKey = os.Getenv("TAXJAR_SANDBOX_API_KEY")
client.APIURL = taxjar.SandboxAPIURL
```

## Testing

```bash
make test
```

To validate API methods in the TaxJar sandbox environment, pass the following environment variables:

```bash
TAXJAR_API_URL="https://api.sandbox.taxjar.com" \
TAXJAR_API_KEY="9e0cd62a22f451701f29c3bde214" \
make test
```

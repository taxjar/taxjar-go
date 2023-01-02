package main

import (
	"fmt"
	"os"

	"github.com/taxjar/taxjar-go"
)

func main() {
	client := taxjar.NewClient(taxjar.Config{
		APIKey: os.Getenv("TAXJAR_API_KEY"),
	})
	// or
	// client := taxjar.NewClient()
	// client.APIKey = os.Getenv("TAXJAR_API_KEY")

	res1, err := client.Categories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nCategories %+v", res1.Categories)

	res2, err := client.TaxForOrder(taxjar.TaxForOrderParams{
		FromCountry: "US",
		FromZip:     "92093",
		FromState:   "CA",
		FromCity:    "La Jolla",
		FromStreet:  "9500 Gilman Drive",
		ToCountry:   "US",
		ToZip:       "90002",
		ToState:     "CA",
		ToCity:      "Los Angeles",
		ToStreet:    "1335 E 103rd St",
		Amount:      15,
		Shipping:    1.5,
		LineItems: []taxjar.TaxLineItem{
			{
				ID:             "1",
				Quantity:       1,
				ProductTaxCode: "20010",
				UnitPrice:      15,
				Discount:       0,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nTaxForOrder %+v", res2.Tax)

	res3, err := client.CreateOrder(taxjar.CreateOrderParams{
		TransactionID:   "13579-246810",
		TransactionDate: "2015/09/08",
		ToCountry:       "US",
		ToZip:           "10019",
		ToState:         "NY",
		ToCity:          "New York",
		ToStreet:        "1697 Broadway",
		Amount:          36.21,
		Shipping:        5,
		SalesTax:        0,
		LineItems: []taxjar.OrderLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "12-34243-9",
				Description:       "Fuzzy Sweater",
				ProductTaxCode:    "20010",
				UnitPrice:         36.72,
				Discount:          5.51,
				SalesTax:          0,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nCreateOrder %+v", res3.Order)

	res4, err := client.ListOrders(taxjar.ListOrdersParams{
		FromTransactionDate: "2015/09/08",
		ToTransactionDate:   "2030/08/19",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nListOrders %+v", res4.Orders)

	res5, err := client.UpdateOrder(taxjar.UpdateOrderParams{
		TransactionID: "13579-246810",
		Amount:        taxjar.Float64(152.72),
		Shipping:      taxjar.Float64(10),
		SalesTax:      taxjar.Float64(10.74),
		LineItems: []taxjar.OrderLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "12-34243-9",
				Description:       "Fuzzy Sweater",
				ProductTaxCode:    "20010",
				UnitPrice:         36.72,
				Discount:          5.51,
				SalesTax:          0,
			},
			{
				ID:                "2",
				Quantity:          1,
				ProductIdentifier: "12-34245-8",
				Description:       "TaxJar Designer T-shirt",
				ProductTaxCode:    "20010",
				UnitPrice:         111,
				SalesTax:          9.85,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nUpdateOrder %+v", res5.Order)

	res6, err := client.ShowOrder("13579-246810")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nShowOrder %+v", res6.Order)

	res7, err := client.DeleteOrder("13579-246810")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nDeleteOrder %+v", res7.Order)

	res8, err := client.CreateRefund(taxjar.CreateRefundParams{
		TransactionID:          "13579-246810-refund",
		TransactionReferenceID: "13579-246810",
		TransactionDate:        "2015/09/08",
		ToCountry:              "US",
		ToZip:                  "10019",
		ToState:                "NY",
		ToCity:                 "New York",
		ToStreet:               "1697 Broadway",
		Amount:                 -116.51,
		Shipping:               -0,
		SalesTax:               -10.74,
		LineItems: []taxjar.RefundLineItem{
			{
				ID:                "1",
				Quantity:          1,
				ProductIdentifier: "12-34243-9",
				Description:       "Fuzzy Sweater",
				ProductTaxCode:    "20010",
				UnitPrice:         -0,
				Discount:          -0,
				SalesTax:          -0,
			},
			{
				ID:                "2",
				Quantity:          1,
				ProductIdentifier: "12-34245-8",
				Description:       "TaxJar Designer T-shirt",
				ProductTaxCode:    "20010",
				UnitPrice:         -111,
				SalesTax:          -9.85,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nCreateRefund %+v", res8.Refund)

	res9, err := client.ListRefunds(taxjar.ListRefundsParams{
		FromTransactionDate: "2015/09/08",
		ToTransactionDate:   "2030/08/19",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nListRefunds %+v", res9.Refunds)

	res10, err := client.UpdateRefund(taxjar.UpdateRefundParams{
		TransactionID:          "13579-246810-refund",
		TransactionReferenceID: "13579-246810",
		Shipping:               -5,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nUpdateRefund %+v", res10.Refund)

	res11, err := client.ShowRefund("13579-246810-refund")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nShowRefund %+v", res11.Refund)

	res12, err := client.DeleteRefund("13579-246810-refund")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nDeleteRefund %+v", res12.Refund)

	res13, err := client.CreateCustomer(taxjar.CreateCustomerParams{
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
		Street:  "4120 Freidrich Ln",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nCreateCustomer %+v", res13.Customer)

	res14, err := client.ListCustomers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nListCustomers %+v", res14.Customers)

	res15, err := client.UpdateCustomer(taxjar.UpdateCustomerParams{
		CustomerID:    "123",
		ExemptionType: "non_exempt",
		Name:          "Initech",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nUpdateCustomer %+v", res15.Customer)

	res16, err := client.ShowCustomer("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nShowCustomer %+v", res16.Customer)

	res17, err := client.DeleteCustomer("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nDeleteCustomer %+v", res17.Customer)

	res18, err := client.RatesForLocation("94043", taxjar.RatesForLocationParams{
		Country: "US",
		State:   "CA",
		City:    "Mountain View",
		Street:  "311 Moffett Blvd",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nRatesForLocation %+v", res18.Rate)

	res19, err := client.NexusRegions()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nNexusRegions %+v", res19.Regions)

	res20, err := client.ValidateAddress(taxjar.ValidateAddressParams{
		Country: "US",
		State:   "AZ",
		Zip:     "85297",
		City:    "Gilbert",
		Street:  "3301 South Greenfield Rd",
	})
	if err := err.(*taxjar.Error); err != nil {
		fmt.Println("\n\nValidateAddress", err)
		// fmt.Println(err.Status)
		// fmt.Println(err.Err)
		// fmt.Println(err.Detail)
		// fmt.Printf("%+v", errors.Wrap(err, ""))
	} else {
		fmt.Printf("\n\nValidateAddress %+v", res20.Addresses)
	}

	res21, err := client.Validate(taxjar.ValidateParams{
		VAT: "FR40303265045",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nValidate %+v", res21.Validation)

	res22, err := client.SummaryRates()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\nSummaryRates %+v", res22.SummaryRates)
}

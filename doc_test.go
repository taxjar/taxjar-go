package taxjar_test

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/taxjar/taxjar-go"
)

var client = taxjar.NewClient()

func ExampleNewClient() {
	// Configure during instantiation
	client := taxjar.NewClient(taxjar.Config{
		APIKey: os.Getenv("TAXJAR_API_KEY"),
	})

	// Or configure after instantiation
	client = taxjar.NewClient()
	client.APIKey = os.Getenv("TAXJAR_API_KEY")
}

func ExampleConfig_Categories() {
	res, err := client.Categories()
	if err != nil {
		// handle error
	}
	fmt.Printf("Categories %+v", res.Categories)
}

func ExampleConfig_TaxForOrder() {
	res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
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
		// handle error
	}
	fmt.Printf("TaxForOrder %+v", res.Tax)
}

func ExampleConfig_CreateOrder() {
	res, err := client.CreateOrder(taxjar.CreateOrderParams{
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
		// handle error
	}
	fmt.Printf("CreateOrder %+v", res.Order)
}

func ExampleConfig_ListOrders() {
	res, err := client.ListOrders(taxjar.ListOrdersParams{
		FromTransactionDate: "2015/09/01",
		ToTransactionDate:   "2015/09/30",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ListOrders %+v", res.Orders)
}

func ExampleConfig_UpdateOrder() {
	res, err := client.UpdateOrder(taxjar.UpdateOrderParams{
		TransactionID: "13579-246810",
		Amount:        152.72,
		Shipping:      10,
		SalesTax:      10.74,
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
		// handle error
	}
	fmt.Printf("UpdateOrder %+v", res.Order)
}

func ExampleConfig_ShowOrder() {
	res, err := client.ShowOrder("13579-246810")
	if err != nil {
		// handle error
	}
	fmt.Printf("ShowOrder %+v", res.Order)
}

func ExampleConfig_DeleteOrder() {
	res, err := client.DeleteOrder("13579-246810")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("DeleteOrder %+v", res.Order)
}

func ExampleConfig_CreateRefund() {
	res, err := client.CreateRefund(taxjar.CreateRefundParams{
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
		// handle error
	}
	fmt.Printf("CreateRefund %+v", res.Refund)
}

func ExampleConfig_ListRefunds() {
	res, err := client.ListRefunds(taxjar.ListRefundsParams{
		FromTransactionDate: "2015/09/01",
		ToTransactionDate:   "2015/09/30",
	})
	if err != nil {
		// handle error
	}
	fmt.Printf("ListRefunds %+v", res.Refunds)
}

func ExampleConfig_UpdateRefund() {
	res, err := client.UpdateRefund(taxjar.UpdateRefundParams{
		TransactionID:          "13579-246810-refund",
		TransactionReferenceID: "13579-246810",
		Shipping:               -5,
	})
	if err != nil {
		// handle error
	}
	fmt.Printf("UpdateRefund %+v", res.Refund)
}

func ExampleConfig_ShowRefund() {
	res, err := client.ShowRefund("13579-246810-refund")
	if err != nil {
		// handle error
	}
	fmt.Printf("ShowRefund %+v", res.Refund)
}

func ExampleConfig_DeleteRefund() {
	res, err := client.DeleteRefund("13579-246810-refund")
	if err != nil {
		// handle error
	}
	fmt.Printf("DeleteRefund %+v", res.Refund)
}

func ExampleConfig_CreateCustomer() {
	res, err := client.CreateCustomer(taxjar.CreateCustomerParams{
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
		// handle error
	}
	fmt.Printf("CreateCustomer %+v", res.Customer)
}

func ExampleConfig_ListCustomers() {
	res, err := client.ListCustomers()
	if err != nil {
		// handle error
	}
	fmt.Printf("ListCustomers %+v", res.Customers)
}

func ExampleConfig_UpdateCustomer() {
	res, err := client.UpdateCustomer(taxjar.UpdateCustomerParams{
		CustomerID:    "123",
		ExemptionType: "non_exempt",
		Name:          "Initech",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("UpdateCustomer %+v", res.Customer)
}

func ExampleConfig_ShowCustomer() {
	res, err := client.ShowCustomer("123")
	if err != nil {
		// handle error
	}
	fmt.Printf("ShowCustomer %+v", res.Customer)
}

func ExampleConfig_DeleteCustomer() {
	res, err := client.DeleteCustomer("123")
	if err != nil {
		// handle error
	}
	fmt.Printf("DeleteCustomer %+v", res.Customer)
}

func ExampleConfig_RatesForLocation() {
	res, err := client.RatesForLocation("94043", taxjar.RatesForLocationParams{
		Country: "US",
		State:   "CA",
		City:    "Mountain View",
		Street:  "311 Moffett Blvd",
	})
	if err != nil {
		// handle error
	}
	fmt.Printf("RatesForLocation %+v", res.Rate)
}

func ExampleConfig_NexusRegions() {
	res, err := client.NexusRegions()
	if err != nil {
		// handle error
	}
	fmt.Printf("NexusRegions %+v", res.Regions)
}

func ExampleConfig_ValidateAddress() {
	res, err := client.ValidateAddress(taxjar.ValidateAddressParams{
		Country: "US",
		State:   "AZ",
		Zip:     "85297",
		City:    "Gilbert",
		Street:  "3301 South Greenfield Rd",
	})
	if err != nil {
		// handle error
	}
	fmt.Printf("ValidateAddress %+v", res.Addresses)
}

func ExampleConfig_Validate() {
	res, err := client.Validate(taxjar.ValidateParams{
		VAT: "FR40303265045",
	})
	if err != nil {
		// handle error
	}
	fmt.Printf("Validate %+v", res.Validation)
}

func ExampleConfig_SummaryRates() {
	res, err := client.SummaryRates()
	if err != nil {
		// handle error
	}
	fmt.Printf("SummaryRates %+v", res.SummaryRates)
}

func ExampleError() {
	client := taxjar.NewClient(taxjar.Config{
		APIKey: "INVALID_API_KEY",
	})
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
		LineItems: []taxjar.TaxLineItem{
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
		fmt.Println(err.Status)                 // 401
		fmt.Println(err.Err)                    // Unauthorized
		fmt.Println(err.Detail)                 // Not authorized for route `POST /v2/taxes'
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
}

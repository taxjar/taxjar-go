package test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/taxjar/taxjar-go"
)

var _ = Describe("using a live/sandbox token", func() {
	RegisterFailHandler(Fail)
	defer GinkgoRecover()

	var client taxjar.Config

	BeforeEach(func() {
		if !IsLiveTestRun {
			Skip("TAXJAR_API_KEY environment variable must be set")
		}
		client = taxjar.NewClient(taxjar.Config{
			APIKey: os.Getenv("TAXJAR_API_KEY"),
			APIURL: os.Getenv("TAXJAR_API_URL"),
		})
	})

	Context("Categories", func() {
		It("lists tax categories", func() {
			res, err := client.Categories()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Categories).To(Not(BeNil()))
		})
	})

	Context("TaxForOrder", func() {
		It("calculates tax for a US-based order", func() {
			res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
				FromCountry:   "US",
				FromZip:       "92093",
				FromState:     "CA",
				FromCity:      "La Jolla",
				FromStreet:    "9500 Gilman Drive",
				ToCountry:     "US",
				ToZip:         "90002",
				ToState:       "CA",
				ToCity:        "Los Angeles",
				ToStreet:      "1335 E 103rd St",
				Amount:        15,
				Shipping:      1.5,
				CustomerID:    "123",
				ExemptionType: "non_exempt",
				NexusAddresses: []taxjar.NexusAddress{
					{
						Country: "US",
						State:   "CA",
					},
				},
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
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Tax).To(Not(BeNil()))
		})
		It("calculates tax for a CA-based order", func() {
			res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
				FromCountry:   "US",
				FromZip:       "92093",
				FromState:     "CA",
				FromCity:      "La Jolla",
				FromStreet:    "9500 Gilman Drive",
				ToCountry:     "CA",
				ToZip:         "V6C 2Y9",
				ToState:       "BC",
				ToCity:        "Vancouver",
				ToStreet:      "645 Howe Street",
				Amount:        15,
				Shipping:      1.5,
				CustomerID:    "123",
				ExemptionType: "non_exempt",
				NexusAddresses: []taxjar.NexusAddress{
					{
						Country: "CA",
						State:   "BC",
					},
				},
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
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Tax).To(Not(BeNil()))
		})
		It("calculates tax for an EU-based order", func() {
			res, err := client.TaxForOrder(taxjar.TaxForOrderParams{
				FromCountry:   "US",
				FromZip:       "92093",
				FromState:     "CA",
				FromCity:      "La Jolla",
				FromStreet:    "9500 Gilman Drive",
				ToCountry:     "FR",
				ToZip:         "69205",
				ToCity:        "Lyon",
				ToStreet:      "1 Place de la Comedie",
				Amount:        15,
				Shipping:      1.5,
				CustomerID:    "123",
				ExemptionType: "non_exempt",
				NexusAddresses: []taxjar.NexusAddress{
					{
						Country: "FR",
					},
				},
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
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Tax).To(Not(BeNil()))
		})
	})

	Context("ListOrders", func() {
		It("lists orders", func() {
			res, err := client.ListOrders(taxjar.ListOrdersParams{
				TransactionDate: "2019/08/26",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Orders).NotTo(BeNil())
		})
	})

	Context("CreateOrder", func() {
		It("creates an order", func() {
			res, err := client.CreateOrder(taxjar.CreateOrderParams{
				TransactionID:   "24",
				TransactionDate: "2019/08/26",
				Provider:        "api",
				FromCountry:     "US",
				FromZip:         "94043",
				FromState:       "CA",
				FromCity:        "Mountain View",
				FromStreet:      "311 Moffett Blvd",
				ToCountry:       "US",
				ToZip:           "10019",
				ToState:         "NY",
				ToCity:          "New York",
				ToStreet:        "1697 Broadway",
				Amount:          50,
				Shipping:        5,
				SalesTax:        0,
				CustomerID:      "123",
				ExemptionType:   "non_exempt",
				LineItems: []taxjar.OrderLineItem{
					{
						ID:                "1",
						Quantity:          1,
						ProductIdentifier: "10-12345-987",
						Description:       "10-gallon Hat",
						ProductTaxCode:    "20010",
						UnitPrice:         45,
						Discount:          0,
						SalesTax:          0,
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Order).NotTo(BeNil())
		})
	})

	Context("ShowOrder", func() {
		It("shows an order", func() {
			res, err := client.ShowOrder("24", taxjar.ShowOrderParams{
				Provider: "api",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Order).NotTo(BeNil())
		})
	})

	Context("UpdateOrder", func() {
		It("updates an order", func() {
			res, err := client.UpdateOrder(taxjar.UpdateOrderParams{
				TransactionID: "24",
				Amount:        161,
				Shipping:      5,
				SalesTax:      10.3,
				LineItems: []taxjar.OrderLineItem{
					{
						ID:                "1",
						Quantity:          1,
						ProductIdentifier: "10-12345-987",
						Description:       "10-gallon Hat",
						ProductTaxCode:    "20010",
						UnitPrice:         45,
						Discount:          0,
						SalesTax:          0,
					},
					{
						ID:                "2",
						Quantity:          1,
						ProductIdentifier: "78-95432-101",
						Description:       "Extra-long Chaps",
						ProductTaxCode:    "20010",
						UnitPrice:         111,
						Discount:          0,
						SalesTax:          9.85,
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Order).NotTo(BeNil())
		})
	})

	Context("DeleteOrder", func() {
		It("deletes an order", func() {
			res, err := client.DeleteOrder("24", taxjar.DeleteOrderParams{Provider: "api"})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Order).NotTo(BeNil())
		})
	})

	Context("ListRefunds", func() {
		It("lists refunds", func() {
			res, err := client.ListRefunds(taxjar.ListRefundsParams{
				TransactionDate: "2019/08/26",
				Provider:        "api",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Refunds).NotTo(BeNil())
		})
	})

	Context("CreateRefund", func() {
		It("creates a refund", func() {
			res, err := client.CreateRefund(taxjar.CreateRefundParams{
				TransactionID:          "24-refund",
				TransactionReferenceID: "24",
				TransactionDate:        "2019/08/26",
				Provider:               "api",
				FromCountry:            "US",
				FromZip:                "94043",
				FromState:              "CA",
				FromCity:               "Mountain View",
				FromStreet:             "311 Moffett Blvd",
				ToCountry:              "US",
				ToZip:                  "10019",
				ToState:                "NY",
				ToCity:                 "New York",
				ToStreet:               "1697 Broadway",
				Amount:                 -111,
				Shipping:               -0,
				SalesTax:               -10.3,
				CustomerID:             "123",
				ExemptionType:          "non_exempt",
				LineItems: []taxjar.RefundLineItem{
					{
						ID:                "1",
						Quantity:          1,
						ProductIdentifier: "10-12345-987",
						Description:       "10-gallon Hat",
						ProductTaxCode:    "20010",
						UnitPrice:         -0,
						Discount:          -0,
						SalesTax:          -0,
					},
					{
						ID:                "2",
						Quantity:          1,
						ProductIdentifier: "78-95432-101",
						Description:       "Extra-long Chaps",
						ProductTaxCode:    "20010",
						UnitPrice:         -111,
						Discount:          -0,
						SalesTax:          -9.85,
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Refund).NotTo(BeNil())
		})
	})

	Context("UpdateRefund", func() {
		It("updates a refund", func() {
			res, err := client.UpdateRefund(taxjar.UpdateRefundParams{
				TransactionID:          "24-refund",
				TransactionReferenceID: "24",
				Amount:                 -116,
				Shipping:               -5,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Refund).NotTo(BeNil())
		})
	})

	Context("DeleteRefund", func() {
		It("deletes a refund", func() {
			res, err := client.DeleteRefund("24-refund", taxjar.DeleteRefundParams{
				Provider: "api",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Refund).NotTo(BeNil())
		})
	})

	Describe("Customer endpoint:", func() {
		BeforeEach(func() {
			if client.APIURL == taxjar.SandboxAPIURL {
				Skip("Customer endpoints not available in sandbox; switch to live token to test")
			}
		})

		Context("ListCustomers", func() {
			It("lists customers", func() {
				res, err := client.ListCustomers()
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Customers).NotTo(BeNil())
			})
		})

		Context("CreateCustomer", func() {
			It("creates a customer", func() {
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
					Street:  "4120 Freidrich Lane",
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Customer).NotTo(BeNil())
			})
		})

		Context("ShowCustomer", func() {
			It("shows a customer", func() {
				res, err := client.ShowCustomer("123")
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Customer).NotTo(BeNil())
			})
		})

		Context("UpdateCustomer", func() {
			It("updates a customer", func() {
				res, err := client.UpdateCustomer(taxjar.UpdateCustomerParams{
					CustomerID:    "123",
					ExemptionType: "non_exempt",
					Name:          "Initech",
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Customer).NotTo(BeNil())
			})
		})

		Context("DeleteCustomer", func() {
			It("deletes a customer", func() {
				res, err := client.DeleteCustomer("123")
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Customer).NotTo(BeNil())
			})
		})
	})

	Context("RatesForLocation", func() {
		It("looks up rates", func() {
			res, err := client.RatesForLocation("89001", taxjar.RatesForLocationParams{
				Country: "US",
				State:   "NV",
				City:    "Alamo",
				Street:  "Mail Box Rd",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rate).NotTo(BeNil())
		})
	})

	Context("NexusRegions", func() {
		It("lists nexus regions", func() {
			res, err := client.NexusRegions()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Regions).NotTo(BeNil())
		})
	})

	// Remove the `X` in `XContext to run ValidateAddress tests`
	XContext("ValidateAddress", func() {
		It("validates an address", func() {
			res, err := client.ValidateAddress(taxjar.ValidateAddressParams{
				Country: "US",
				State:   "AZ",
				Zip:     "85297",
				City:    "Gilbert",
				Street:  "3301 South Greenfield Rd",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Addresses).NotTo(BeNil())
		})
	})

	Context("Validate", func() {
		It("validates an VAT identification number", func() {
			res, err := client.Validate(taxjar.ValidateParams{
				VAT: "FR40303265045",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Validation).NotTo(BeNil())
		})
	})

	Context("SummaryRates", func() {
		It("summarizes rates", func() {
			res, err := client.SummaryRates()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.SummaryRates).NotTo(BeNil())
		})
	})

})

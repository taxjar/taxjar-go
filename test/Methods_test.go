package test

import (
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/taxjar/taxjar-go"
	"github.com/taxjar/taxjar-go/test/mocks"
)

var IsLiveTestRun = os.Getenv("TAXJAR_API_KEY") != ""

var _ = Describe("Method:", func() {
	RegisterFailHandler(Fail)
	defer GinkgoRecover()

	var server *ghttp.Server
	var client taxjar.Config

	BeforeEach(func() {
		if IsLiveTestRun {
			Skip("TAXJAR_API_KEY environment variable must not be set")
		}
		server = ghttp.NewServer()
		client = taxjar.NewClient(taxjar.Config{
			APIKey: "test123",
			APIURL: server.URL(),
		})
	})

	AfterEach(func() {
		server.Close()
	})

	Context("Error", func() {
		It("s/b an error string, and allow extracting `Status`, `Err`, & `Detail`", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/categories"),
				ghttp.RespondWith(401, mocks.ErrorJSON),
			))
			_, err := client.Categories()
			Expect(err).To(MatchError(mocks.Error.Error()))
			// assert to `*taxjar.Error` to extract details
			if err := err.(*taxjar.Error); true {
				Expect(err.Status).To(Equal(mocks.Error.Status))
				Expect(err.Err).To(Equal(mocks.Error.Err))
				Expect(err.Detail).To(Equal(mocks.Error.Detail))
			}
		})
	})

	Context("Categories", func() {
		It("lists tax categories", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/categories"),
				ghttp.RespondWith(http.StatusOK, mocks.CategoriesJSON),
			))
			res, err := client.Categories()
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(&mocks.Categories))
		})
	})

	Context("TaxForOrder", func() {
		It("calculates tax for a US-based order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/taxes"),
				ghttp.RespondWith(http.StatusOK, mocks.USTaxForOrderJSON),
			))
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.USTaxForOrder))
		})
		It("calculates tax for a CA-based order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/taxes"),
				ghttp.RespondWith(http.StatusOK, mocks.CATaxForOrderJSON),
			))
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.CATaxForOrder))
		})
		It("calculates tax for an EU-based order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/taxes"),
				ghttp.RespondWith(http.StatusOK, mocks.EUTaxForOrderJSON),
			))
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.EUTaxForOrder))
		})
	})

	Context("ListOrders", func() {
		It("lists orders", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/transactions/orders", "transaction_date=2019%2F08%2F26"),
				ghttp.RespondWith(http.StatusOK, mocks.ListOrdersJSON),
			))
			res, err := client.ListOrders(taxjar.ListOrdersParams{
				TransactionDate: "2019/08/26",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ListOrders))
		})
	})

	Context("ShowOrder", func() {
		It("shows an order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/transactions/orders/24"),
				ghttp.RespondWith(http.StatusOK, mocks.ShowOrderJSON),
			))
			res, err := client.ShowOrder("24")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ShowOrder))
		})
		It("shows an order with params", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/transactions/orders/24", "provider=api"),
				ghttp.RespondWith(http.StatusOK, mocks.ShowOrderJSON),
			))
			res, err := client.ShowOrder("24", taxjar.ShowOrderParams{
				Provider: "api",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ShowOrder))
		})
	})

	Context("CreateOrder", func() {
		It("creates an order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/transactions/orders"),
				ghttp.RespondWith(http.StatusOK, mocks.CreateOrderJSON),
			))
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.CreateOrder))
		})
	})

	Context("UpdateOrder", func() {
		It("updates an order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("PUT", "/v2/transactions/orders/24"),
				ghttp.RespondWith(http.StatusOK, mocks.UpdateOrderJSON),
			))

			res, err := client.UpdateOrder(taxjar.UpdateOrderParams{

				TransactionID: "24",
				Amount:        161,
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.UpdateOrder))
		})
	})

	Context("DeleteOrder", func() {
		It("deletes an order", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("DELETE", "/v2/transactions/orders/24"),
				ghttp.RespondWith(http.StatusOK, mocks.DeleteOrderJSON),
			))
			res, err := client.DeleteOrder("24")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.DeleteOrder))
		})
		It("deletes an order with params", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("DELETE", "/v2/transactions/orders/24", "provider=api"),
				ghttp.RespondWith(http.StatusOK, mocks.DeleteOrderJSON),
			))
			res, err := client.DeleteOrder("24", taxjar.DeleteOrderParams{Provider: "api"})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.DeleteOrder))
		})
	})

	Context("ListRefunds", func() {
		It("lists refunds", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/transactions/refunds", "transaction_date=2019%2F08%2F26&provider=api"),
				ghttp.RespondWith(http.StatusOK, mocks.ListRefundsJSON),
			))
			res, err := client.ListRefunds(taxjar.ListRefundsParams{
				TransactionDate: "2019/08/26",
				Provider:        "api",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ListRefunds))
		})
	})

	Context("ShowRefund", func() {
		It("shows a refund", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/transactions/refunds/24-refund"),
				ghttp.RespondWith(http.StatusOK, mocks.ShowRefundJSON),
			))
			res, err := client.ShowRefund("24-refund")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ShowRefund))
		})
		It("shows a refund with params", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/transactions/refunds/24-refund", "provider=api"),
				ghttp.RespondWith(http.StatusOK, mocks.ShowRefundJSON),
			))
			res, err := client.ShowRefund("24-refund", taxjar.ShowRefundParams{
				Provider: "api",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ShowRefund))
		})
	})

	Context("CreateRefund", func() {
		It("creates a refund", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/transactions/refunds"),
				ghttp.RespondWith(http.StatusOK, mocks.CreateRefundJSON),
			))
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.CreateRefund))
		})
	})

	Context("UpdateRefund", func() {
		It("updates a refund", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("PUT", "/v2/transactions/refunds/24-refund"),
				ghttp.RespondWith(http.StatusOK, mocks.UpdateRefundJSON),
			))
			res, err := client.UpdateRefund(taxjar.UpdateRefundParams{
				TransactionID:          "24-refund",
				TransactionReferenceID: "24",
				Shipping:               -5,
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.UpdateRefund))
		})
	})

	Context("DeleteRefund", func() {
		It("deletes a refund", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("DELETE", "/v2/transactions/refunds/24-refund"),
				ghttp.RespondWith(http.StatusOK, mocks.DeleteRefundJSON),
			))
			res, err := client.DeleteRefund("24-refund")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.DeleteRefund))
		})
		It("deletes a refund with params", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("DELETE", "/v2/transactions/refunds/24-refund", "provider=api"),
				ghttp.RespondWith(http.StatusOK, mocks.DeleteRefundJSON),
			))
			res, err := client.DeleteRefund("24-refund", taxjar.DeleteRefundParams{
				Provider: "api",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.DeleteRefund))
		})
	})

	Context("ListCustomers", func() {
		It("lists customers", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/customers"),
				ghttp.RespondWith(http.StatusOK, mocks.ListCustomersJSON),
			))
			res, err := client.ListCustomers()
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ListCustomers))
		})
	})

	Context("ShowCustomer", func() {
		It("shows a customer", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/customers/123"),
				ghttp.RespondWith(http.StatusOK, mocks.ShowCustomerJSON),
			))
			res, err := client.ShowCustomer("123")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ShowCustomer))
		})
	})

	Context("CreateCustomer", func() {
		It("creates a customer", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/customers"),
				ghttp.RespondWith(http.StatusOK, mocks.CreateCustomerJSON),
			))
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
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.CreateCustomer))
		})
	})

	Context("UpdateCustomer", func() {
		It("updates a customer", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("PUT", "/v2/customers/123"),
				ghttp.RespondWith(http.StatusOK, mocks.UpdateCustomerJSON),
			))
			res, err := client.UpdateCustomer(taxjar.UpdateCustomerParams{
				CustomerID:    "123",
				ExemptionType: "non_exempt",
				Name:          "Initech",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.UpdateCustomer))
		})
	})

	Context("DeleteCustomer", func() {
		It("deletes a customer", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("DELETE", "/v2/customers/123"),
				ghttp.RespondWith(http.StatusOK, mocks.DeleteCustomerJSON),
			))
			res, err := client.DeleteCustomer("123")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.DeleteCustomer))
		})
	})

	Context("RatesForLocation", func() {
		It("looks up rates", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/rates/89001"),
				ghttp.RespondWith(http.StatusOK, mocks.RatesForLocationJSON),
			))
			res, err := client.RatesForLocation("89001")
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.RatesForLocation))
		})
		It("looks up rates with params", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/rates/89001", "country=US&state=NV&city=Alamo&street=Mail%20Box%20Rd"),
				ghttp.RespondWith(http.StatusOK, mocks.RatesForLocationJSON),
			))
			res, err := client.RatesForLocation("89001", taxjar.RatesForLocationParams{
				Country: "US",
				State:   "NV",
				City:    "Alamo",
				Street:  "Mail Box Rd",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.RatesForLocation))
		})
	})

	Context("NexusRegions", func() {
		It("lists nexus regions", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/nexus/regions"),
				ghttp.RespondWith(http.StatusOK, mocks.NexusRegionsJSON),
			))
			res, err := client.NexusRegions()
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.NexusRegions))
		})
	})

	Context("ValidateAddress", func() {
		It("validates an address", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/v2/addresses/validate"),
				ghttp.RespondWith(http.StatusOK, mocks.ValidateAddressJSON),
			))
			res, err := client.ValidateAddress(taxjar.ValidateAddressParams{
				Country: "US",
				State:   "AZ",
				Zip:     "85297",
				City:    "Gilbert",
				Street:  "3301 South Greenfield Rd",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.ValidateAddress))
		})
	})

	Context("Validate", func() {
		It("validates a VAT identification number", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/validation", "vat=FR40303265045"),
				ghttp.RespondWith(http.StatusOK, mocks.ValidateJSON),
			))
			res, err := client.Validate(taxjar.ValidateParams{
				VAT: "FR40303265045",
			})
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.Validate))
		})
	})

	Context("SummaryRates", func() {
		It("summarizes rates", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v2/summary_rates"),
				ghttp.RespondWith(http.StatusOK, mocks.SummaryRatesJSON),
			))
			res, err := client.SummaryRates()
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(mocks.SummaryRates))
		})
	})

})

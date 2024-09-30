package payment_gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com/xendit/xendit-go/v6"
	"github.com/xendit/xendit-go/v6/invoice"
	"log"
	"strconv"
)

type XenditClient struct {
	client     *xendit.APIClient
	successUrl string
	failureUrl string
}

type IPaymentGateway interface {
	CreateInvoice(ctx context.Context, order entity.Order) (*invoice.Invoice, error)
}

func NewXendit(cfg config.PaymentGatewayConfig) IPaymentGateway {
	client := xendit.NewClient(cfg.SecretKey)
	return &XenditClient{
		client:     client,
		successUrl: cfg.SuccessRedirectUrl,
		failureUrl: cfg.FailureRedirectUrl,
	}
}

func (x *XenditClient) CreateInvoice(ctx context.Context, order entity.Order) (*invoice.Invoice, error) {
	item := invoice.InvoiceItem{
		Name:     order.Product.Name,
		Price:    float32(order.Product.Price),
		Quantity: float32(order.Quantity),
		Category: &order.Product.Category,
	}
	buyerIDStr := strconv.Itoa(order.Buyer.ID)

	duration := "86400"
	log.Println(order.GrandTotal)
	createInvoiceReq := *invoice.NewCreateInvoiceRequest("example_id", float64(order.GrandTotal))
	createInvoiceReq.Items = append(createInvoiceReq.Items, item)
	createInvoiceReq.Customer = &invoice.CustomerObject{
		CustomerId:  *invoice.NewNullableString(&buyerIDStr),
		GivenNames:  *invoice.NewNullableString(&order.Buyer.Name),
		Email:       *invoice.NewNullableString(&order.Buyer.Email),
		PhoneNumber: *invoice.NewNullableString(&order.Buyer.PhoneNumber),
	}

	createInvoiceReq.SuccessRedirectUrl = &x.successUrl
	createInvoiceReq.FailureRedirectUrl = &x.failureUrl

	createInvoiceReq.InvoiceDuration = &duration
	createInvoiceReq.Fees = []invoice.InvoiceFee{
		{"Platform fee", 1000},
	}
	invoiceResp, httpResp, errXendit := x.client.InvoiceApi.CreateInvoice(ctx).CreateInvoiceRequest(createInvoiceReq).Execute()

	if errXendit != nil {
		b, _ := json.Marshal(errXendit.FullError())
		fmt.Printf("Error when try to get balance with error detail : %v\n", string(b))
		fmt.Printf("Full HTTP response: %v\n", httpResp)
		return invoiceResp, errXendit
	}

	if invoiceResp == nil {
		newErr := errors.New("invoice xendit is nil")
		fmt.Printf("Full HTTP response: %v\n", httpResp)
		fmt.Printf("Error when try to get balance with error detail : %v\n", newErr.Error())
		return invoiceResp, newErr
	}

	return invoiceResp, nil

}

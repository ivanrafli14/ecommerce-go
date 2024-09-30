package entity

import "time"

type CheckoutRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Order struct {
	ID          string        `json:"id" bson:"id"`
	InvoiceID   string        `json:"-" bson:"invoice_id"`
	Quantity    int           `json:"quantity" bson:"quantity"`
	Price       int           `json:"price" bson:"price"`
	SubTotal    int           `json:"sub_total" bson:"sub_total"`
	PlatformFee int           `json:"platform_fee" bson:"platform_fee"`
	GrandTotal  int           `json:"grand_total" bson:"grand_total"`
	InvoiceUrl  string        `json:"invoice_url" bson:"invoice_url"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	Status      string        `json:"status" bson:"status"`
	Buyer       UserBuyer     `json:"-" bson:"buyer"`
	Product     ProductOrder  `json:"product" bson:"product"`
	Merchant    MerchantOrder `json:"merchant" bson:"merchant"`
}

type ProductOrder struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	Stock       int    `json:"stock" bson:"stock"`
	Category    string `json:"category" bson:"category"`
	ImageUrl    string `json:"image_url" bson:"image_url"`
}

type MerchantOrder struct {
	ID       int    `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	ImageUrl string `json:"image_url" bson:"image_url"`
}

type MeilisearchOrderResponse struct {
	ID          string       `json:"id"`
	Quantity    int          `json:"quantity"`
	Price       int          `json:"price"`
	SubTotal    int          `json:"sub_total"`
	PlatformFee int          `json:"platform_fee"`
	GrandTotal  int          `json:"grand_total"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Status      string       `json:"status"`
	Product     ProductOrder `json:"product"`
}

type WebhookInvoiceRequest struct {
	// invoice id that xendit generated
	Id string `json:"id"`
	// invoice id that our server generated
	ExternalId string `json:"external_id"`
	// our xendit user id
	UserId string `json:"user_id"`
	IsHigh bool   `json:"is_high"`
	// status if our invoice is PAID or EXPIRED
	Status string `json:"status"`
	// our xendit merchant name
	MerchantName string `json:"merchant_name"`
	// nominal amount for the invoice
	Amount float64 `json:"amount"`
	// total amount paid for the invoice
	PaidAmount float64 `json:"paid_amount"`
	// email user
	PayerEmail string `json:"payer_email"`
	// description for the invoice
	Description string `json:"description"`
	// invoice when updated
	UpdatedAt time.Time `json:"updated_at"`
	// invoice when created
	CreatedAt time.Time `json:"created_at"`
	// invoice when paid
	PaidAt   time.Time `json:"paid_at"`
	Currency string    `json:"currency"`
	// payment channel like BANK, eWallets
	PaymentChannel string `json:"payment_channel"`
	// payment method like DANA, BCA, BRI, OVO
	PaymentMethod      string `json:"payment_method"`
	PaymentDestination string `json:"payment_destination"`
	PaymentId          string `json:"payment_id"`
}

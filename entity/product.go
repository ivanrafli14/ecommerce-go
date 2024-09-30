package entity

import "time"

type Product struct {
	ID          int       `db:"id"`
	SKU         string    `db:"sku"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	Stock       int       `db:"stock"`
	CategoryID  int       `db:"category_id"`
	ImageUrl    string    `db:"image_url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`

	MerchantID int `db:"merchant_id"`
}

type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryID  int    `json:"category_id"`
	ImageUrl    string `json:"image_url"`

	AuthID     int    `json:"-"`
	Role       string `json:"-"`
	MerchantID int    `json:"-"`
}

type ProductDetailResponse struct {
	ID          int       `json:"id" db:"id"`
	SKU         string    `json:"sku" db:"sku"`
	Name        string    `json:"name" db:"name"` // This matches `p.name`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	Stock       int       `json:"stock" db:"stock"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Category    string    `json:"category" db:"category"` // This matches `categories.name AS category`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	MerchantID int `json:"-" db:"merchant_id"`
}

type DetailProductSKUResponse struct {
	ID           int            `json:"id" db:"id"`
	SKU          string         `json:"sku" db:"sku"`
	Name         string         `json:"name" db:"name"`
	Description  string         `json:"description" db:"description"`
	Price        int            `json:"price" db:"price"`
	Stock        int            `json:"stock" db:"stock"`
	CategoryName string         `json:"category" db:"category_name"`
	CategoryID   int            `json:"category_id" db:"category_id"`
	Merchant     MerchantDetail `json:"merchant" `
	ImageUrl     string         `json:"image_url" db:"image_url"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
}

type MerchantDetail struct {
	ID   int    `json:"id" db:"merchant_id"`
	Name string `json:"name" db:"merchant_name"`
	City string `json:"city" db:"merchant_city"`
}

type MeilisearchPayloadResponse struct {
	ID          int    `json:"id"`
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
}

type MeilisearchPagination struct {
	Query     string `json:"query,omitempty"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	TotalPage int    `json:"total_page"`
}

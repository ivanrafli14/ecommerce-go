package meilisearch

import (
	"fmt"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com/meilisearch/meilisearch-go"
	"log"
	"time"
)

type SearchEngine interface {
	SearchQuery(query string, limit int, page int, merchantID *int) ([]entity.MeilisearchPayloadResponse, entity.MeilisearchPagination, error)
	StoreData(product entity.Product, categoryName string) error
	UpdateData(product entity.Product, categoryName string, productID int) error

	StoreDataOrder(order entity.Order) error
	UpdateDataOrder(webhookReq entity.WebhookInvoiceRequest) error
	SearchQueryOrder(limit int, page int, ID int, role string) ([]entity.Order, entity.MeilisearchPagination, error)
}

type MeilisearchClient struct {
	Client meilisearch.ServiceManager
}

func NewMeilisearch(cfg config.MeiliSearch) SearchEngine {

	client := meilisearch.New(cfg.Username, meilisearch.WithAPIKey(cfg.APIKey))
	_, err := client.Index("products").UpdateFilterableAttributes(&[]string{"products", "merchant_id"})
	if err != nil {
		panic(err)
	}
	_, err = client.Index("orders").UpdateFilterableAttributes(&[]string{"orders", "merchant.merchant_id", "buyer.buyer_id"})
	if err != nil {
		panic(err)
	}
	return &MeilisearchClient{
		Client: client,
	}
}
func (m *MeilisearchClient) StoreData(product entity.Product, categoryName string) error {
	log.Println(product.ID)
	document := map[string]interface{}{
		"id":          product.ID,
		"sku":         product.SKU,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"category":    categoryName,
		"image_url":   product.ImageUrl,
		"merchant_id": product.MerchantID,
	}

	_, err := m.Client.Index("products").AddDocuments(document, "id")
	if err != nil {
		return err
	}
	return nil
}

func (m *MeilisearchClient) SearchQuery(query string, limit int, page int, merchantID *int) ([]entity.MeilisearchPayloadResponse, entity.MeilisearchPagination, error) {
	var products []entity.MeilisearchPayloadResponse
	var pagination entity.MeilisearchPagination

	filter := ""
	if merchantID != nil {
		filter = fmt.Sprintf("merchant_id=%d", *merchantID)
	}
	index := m.Client.Index("products")
	searchResult, err := index.Search(query, &meilisearch.SearchRequest{
		HitsPerPage: 100,
		Page:        int64(page),
		Limit:       int64(limit),
		Facets:      []string{"products"},
		Filter:      filter,
	})

	if err != nil {
		return products, pagination, err
	}

	for _, item := range searchResult.Hits {
		itemMap := item.(map[string]interface{})
		product := entity.MeilisearchPayloadResponse{
			ID:          int(itemMap["id"].(float64)),
			SKU:         itemMap["sku"].(string),
			Name:        itemMap["name"].(string),
			Description: itemMap["description"].(string),
			Price:       int(itemMap["price"].(float64)),
			Stock:       int(itemMap["stock"].(float64)),
			Category:    itemMap["category"].(string),
			ImageUrl:    itemMap["image_url"].(string),
		}
		products = append(products, product)

	}

	pagination.Query = query
	pagination.Limit = limit
	pagination.Page = page
	pagination.TotalPage = int(searchResult.TotalPages)

	return products, pagination, nil
}

func (m *MeilisearchClient) UpdateData(product entity.Product, categoryName string, productID int) error {
	document := map[string]interface{}{
		"id":          productID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"category":    categoryName,
		"image_url":   product.ImageUrl,
	}

	_, err := m.Client.Index("products").UpdateDocuments(document, "id")
	if err != nil {
		return err
	}
	return nil
}

func (m *MeilisearchClient) StoreDataOrder(order entity.Order) error {
	document := map[string]interface{}{
		"id":           order.ID,
		"invoice_id":   order.InvoiceID,
		"quantity":     order.Quantity,
		"price":        order.Price,
		"sub_total":    order.SubTotal,
		"platform_fee": order.PlatformFee,
		"grand_total":  order.GrandTotal,
		"created_at":   order.CreatedAt,
		"updated_at":   order.UpdatedAt,
		"status":       order.Status,
		"invoice_url":  order.InvoiceUrl,
		"product": map[string]any{
			"product_id":          order.Product.ID,
			"product_name":        order.Product.Name,
			"product_description": order.Product.Description,
			"product_price":       order.Product.Price,
			"product_stock":       order.Product.Stock,
			"product_category":    order.Product.Category,
			"product_image_url":   order.Product.ImageUrl,
		},
		"merchant": map[string]any{
			"merchant_id":        order.Merchant.ID,
			"merchant_name":      order.Merchant.Name,
			"merchant_image_url": order.Merchant.ImageUrl,
		},
		"buyer": map[string]any{
			"buyer_id":           order.Buyer.ID,
			"buyer_name":         order.Buyer.Name,
			"buyer_email":        order.Buyer.Email,
			"buyer_phone_number": order.Buyer.PhoneNumber,
		},
	}

	_, err := m.Client.Index("orders").AddDocuments(document, "invoice_id")
	if err != nil {
		return err
	}
	return nil
}

func (m *MeilisearchClient) UpdateDataOrder(webhookReq entity.WebhookInvoiceRequest) error {
	document := map[string]interface{}{
		"invoice_id": webhookReq.Id,
		"status":     webhookReq.Status,
	}
	_, err := m.Client.Index("orders").UpdateDocuments(document, "invoice_id")
	if err != nil {
		return err
	}
	return nil
}

func (m *MeilisearchClient) SearchQueryOrder(limit int, page int, ID int, role string) ([]entity.Order, entity.MeilisearchPagination, error) {

	var orders []entity.Order
	var pagination entity.MeilisearchPagination

	index := m.Client.Index("orders")
	filter := ""

	if role == "user" {
		filter = fmt.Sprintf("buyer.buyer_id=%d", ID)
	} else {
		filter = fmt.Sprintf("merchant.merchant_id=%d", ID)
	}
	searchResult, err := index.Search("", &meilisearch.SearchRequest{
		HitsPerPage: 100,
		Page:        int64(page),
		Limit:       int64(limit),
		Facets:      []string{"orders"},
		Filter:      filter,
	})
	if err != nil {
		return orders, pagination, err
	}

	for _, item := range searchResult.Hits {
		itemMap := item.(map[string]interface{})
		order := entity.Order{
			ID:          itemMap["id"].(string),
			Quantity:    int(itemMap["quantity"].(float64)),
			Price:       int(itemMap["price"].(float64)),
			SubTotal:    int(itemMap["sub_total"].(float64)),
			PlatformFee: int(itemMap["platform_fee"].(float64)),
			GrandTotal:  int(itemMap["grand_total"].(float64)),
			CreatedAt:   parseTime(itemMap["created_at"]),
			UpdatedAt:   parseTime(itemMap["updated_at"]),
			Status:      itemMap["status"].(string),
			InvoiceUrl:  itemMap["invoice_url"].(string),
		}
		if productMap, ok := itemMap["product"].(map[string]interface{}); ok {
			order.Product.ID = int(productMap["product_id"].(float64))
			order.Product.Name = productMap["product_name"].(string)
			order.Product.Description = productMap["product_description"].(string)
			order.Product.Price = int(productMap["product_price"].(float64))
			order.Product.Stock = int(productMap["product_stock"].(float64))
			order.Product.Category = productMap["product_category"].(string)
			order.Product.ImageUrl = productMap["product_image_url"].(string)
		}

		if merchantMap, ok := itemMap["merchant"].(map[string]interface{}); ok {
			order.Merchant.ID = int(merchantMap["merchant_id"].(float64))
			order.Merchant.Name = merchantMap["merchant_name"].(string)
			order.Merchant.ImageUrl = merchantMap["merchant_image_url"].(string)
		}

		orders = append(orders, order)

	}

	pagination.Limit = limit
	pagination.Page = page
	pagination.TotalPage = int(searchResult.TotalPages)

	return orders, pagination, nil
}

func parseTime(value interface{}) time.Time {
	switch v := value.(type) {
	case string:
		// Assuming the time is in RFC3339 format, adjust if needed
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			log.Printf(err.Error())
			return time.Time{} // Return zero value if parsing fails
		}
		return parsedTime
	case time.Time:
		return v
	default:
		log.Printf("Unexpected type for date: %T", value)
		return time.Time{} // Return zero value if type is not expected
	}
}

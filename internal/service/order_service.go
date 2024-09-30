package service

import (
	"context"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/meilisearch"
	"github.com.ivanrafli14/ecommerce-golang/pkg/payment_gateway"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com.ivanrafli14/ecommerce-golang/pkg/validation"
	"github.com/google/uuid"
	"time"
)

type IOrderService interface {
	Checkout(checkoutReq entity.CheckoutRequest, authID int) (string, error)
	ListOrdersMerchant(limit, page, authID int) ([]entity.MeilisearchOrderResponse, entity.MeilisearchPagination, error)
	ListOrdersUser(limit, page, authID int) ([]entity.Order, entity.MeilisearchPagination, error)
	WebhookOrder(webhookReq entity.WebhookInvoiceRequest) error
}

type OrderService struct {
	orderRepo      repository.IOrderRepository
	merchantRepo   repository.IMerchantRepository
	productRepo    repository.IProductRepository
	authRepo       repository.IAuthRepository
	userRepo       repository.IUserRepository
	paymentGateway payment_gateway.IPaymentGateway
	meilisearch    meilisearch.SearchEngine
}

func NewOrderService(or repository.IOrderRepository, mr repository.IMerchantRepository, pr repository.IProductRepository, ar repository.IAuthRepository, ur repository.IUserRepository, pg payment_gateway.IPaymentGateway, ms meilisearch.SearchEngine) IOrderService {
	return &OrderService{
		orderRepo:      or,
		merchantRepo:   mr,
		productRepo:    pr,
		authRepo:       ar,
		userRepo:       ur,
		paymentGateway: pg,
		meilisearch:    ms,
	}
}

func (s *OrderService) Checkout(checkoutReq entity.CheckoutRequest, authID int) (string, error) {
	var order entity.Order

	if err := validation.ValidateCheckoutReq(checkoutReq); err != nil {
		return "", err
	}
	product, err := s.productRepo.GetProductByProductID(checkoutReq.ProductID, -1)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return "", response.ErrProductOrderNotFound
		}
		return "", err
	}

	if product.Stock-checkoutReq.Quantity < 0 {
		return "", response.ErrMaxReachQuantity
	}

	merchant, err := s.merchantRepo.GetMerchantByID(product.MerchantID)
	if err != nil {
		return "", err
	}

	auth, err := s.authRepo.GetAuthByID(authID)
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.FindByAuthID(authID)
	if err != nil {
		return "", err
	}

	order.ID = uuid.New().String()
	order.Quantity = checkoutReq.Quantity
	order.Price = product.Price
	order.SubTotal = order.Price * order.Quantity
	order.PlatformFee = 1000
	order.GrandTotal = order.PlatformFee + order.SubTotal
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Status = "UNPAID"

	order.Product.ID = product.ID
	order.Product.Name = product.Name
	order.Product.Description = product.Description
	order.Product.Price = product.Price
	order.Product.Stock = product.Stock
	order.Product.Category = product.Category
	order.Product.ImageUrl = product.ImageUrl

	order.Merchant.ID = merchant.ID
	order.Merchant.Name = merchant.Name
	order.Merchant.ImageUrl = merchant.ImageUrl

	order.Buyer = entity.UserBuyer{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       auth.Email,
	}

	invoiceResp, err := s.paymentGateway.CreateInvoice(context.Background(), order)
	if err != nil {
		return "", err
	}
	order.InvoiceUrl = invoiceResp.InvoiceUrl
	order.InvoiceID = *invoiceResp.Id

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return "", err
	}

	if err := s.meilisearch.StoreDataOrder(order); err != nil {

		return "", err
	}

	return order.InvoiceUrl, nil

}

func (s *OrderService) ListOrdersMerchant(limit, page, authID int) ([]entity.MeilisearchOrderResponse, entity.MeilisearchPagination, error) {
	merchant, err := s.merchantRepo.GetMerchantByAuthId(authID)
	if err != nil {
		return nil, entity.MeilisearchPagination{}, err
	}

	orders, pagination, err := s.meilisearch.SearchQueryOrder(limit, page, merchant.ID, "merchant")
	ordersResponse := make([]entity.MeilisearchOrderResponse, len(orders))
	if err != nil {
		return ordersResponse, pagination, err
	}

	for i, order := range orders {
		ordersResponse[i] = entity.MeilisearchOrderResponse{
			ID:          order.ID,
			Quantity:    order.Quantity,
			Price:       order.Price,
			SubTotal:    order.SubTotal,
			PlatformFee: order.PlatformFee,
			GrandTotal:  order.GrandTotal,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
			Status:      order.Status,
		}
		ordersResponse[i].Product.ID = order.Product.ID
		ordersResponse[i].Product.Name = order.Product.Name
		ordersResponse[i].Product.Description = order.Product.Description
		ordersResponse[i].Product.Price = order.Product.Price
		ordersResponse[i].Product.Stock = order.Product.Stock
		ordersResponse[i].Product.Category = order.Product.Category
		ordersResponse[i].Product.ImageUrl = order.Product.ImageUrl
	}
	return ordersResponse, pagination, nil
}

func (s *OrderService) ListOrdersUser(limit, page, authID int) ([]entity.Order, entity.MeilisearchPagination, error) {
	user, err := s.userRepo.FindByAuthID(authID)
	if err != nil {
		return nil, entity.MeilisearchPagination{}, err
	}
	orders, pagination, err := s.meilisearch.SearchQueryOrder(limit, page, user.ID, "user")
	if err != nil {
		return orders, pagination, err
	}
	return orders, pagination, nil
}

func (s *OrderService) WebhookOrder(webhookReq entity.WebhookInvoiceRequest) error {
	err := s.orderRepo.WebhookOrder(webhookReq)
	if err != nil {
		return err
	}

	err = s.meilisearch.UpdateDataOrder(webhookReq)
	if err != nil {
		return err
	}
	return nil

}

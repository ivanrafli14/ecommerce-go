package repository

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/database/mongodb"
)

type IOrderRepository interface {
	CreateOrder(order entity.Order) error
	//UpdateOrder(order entity.Order) error
	WebhookOrder(webhookReq entity.WebhookInvoiceRequest) error
}

type OrderRepository struct {
	db mongodb.IMongoDB
}

func NewOrderRepository(db mongodb.IMongoDB) IOrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order entity.Order) error {
	return r.db.StoreData(order)
}

//func (r *OrderRepository) UpdateOrder(order entity.Order) error {
//	return r.db.UpdateData(order)
//}

func (r *OrderRepository) WebhookOrder(webhookReq entity.WebhookInvoiceRequest) error {
	return r.db.UpdateData(webhookReq)
}

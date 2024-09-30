package mongodb

import (
	"context"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type IMongoDB interface {
	StoreData(order entity.Order) error
	UpdateData(webhookReq entity.WebhookInvoiceRequest) error
	//GetOrders(authID int) ([]entity.Order, error)
}

type MongoDBClient struct {
	mongoDBClient  *mongo.Client
	databaseName   string
	collectionName string
}

func NewMongoDBClient(cfg config.MongoDBConfig) IMongoDB {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	return &MongoDBClient{
		mongoDBClient:  client,
		databaseName:   cfg.DatabaseName,
		collectionName: cfg.CollectionName,
	}
}

func (m *MongoDBClient) StoreData(order entity.Order) error {
	collection := m.mongoDBClient.Database(m.databaseName).Collection(m.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the order into the collection
	_, err := collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	// Set the inserted ID to the order (MongoDB returns _id field)
	//order.ID = result.InsertedID.(string)
	return nil
}

func (m *MongoDBClient) UpdateData(webhookReq entity.WebhookInvoiceRequest) error {
	collection := m.mongoDBClient.Database(m.databaseName).Collection(m.collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println(webhookReq.UserId)
	filter := bson.M{"invoice_id": webhookReq.Id}
	update := bson.M{
		"$set": bson.M{
			"status": webhookReq.Status,
		},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	log.Println(res)
	if err != nil {
		return err
	}
	return nil
}

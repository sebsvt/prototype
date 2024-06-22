package ordering

import (
	"context"
	"time"

	"github.com/google/uuid"
	domain "github.com/sebsvt/prototype/domain/ordering"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepositoryMongoDB struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewOrderRepositoryMongoDB(ctx context.Context, connection_string string) (*OrderRepositoryMongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection_string))
	if err != nil {
		return nil, err
	}
	db := client.Database("prototype")
	orders := db.Collection("orders")
	return &OrderRepositoryMongoDB{
		db:         db,
		collection: orders,
	}, nil
}

func (repo *OrderRepositoryMongoDB) FromReference(order_ref uuid.UUID) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var order domain.Order
	result := repo.collection.FindOne(ctx, bson.M{"reference": order_ref})
	err := result.Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
func (repo *OrderRepositoryMongoDB) FromCustomerID(customer_id uuid.UUID) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var orders []domain.Order
	curs, err := repo.collection.Find(ctx, bson.M{"customer_id": customer_id})
	for curs.Next(context.TODO()) {
		var order domain.Order
		err := curs.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (repo *OrderRepositoryMongoDB) Save(order domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{}
	opts := options.Update().SetUpsert(true)
	update := bson.M{
		"$set": order,
	}
	_, err := repo.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

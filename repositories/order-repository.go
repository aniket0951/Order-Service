package repositories

import (
	"context"
	"os"

	"github.com/aniket0951/order-services/config"
	"github.com/aniket0951/order-services/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ordersConnection       = config.GetCollection(config.DB, os.Getenv("ORDERS"))
	orderCartConnection    = config.GetCollection(config.DB, os.Getenv("ORDER_CART"))
	orderTrackConnection   = config.GetCollection(config.DB, os.Getenv("ORDER_TRACK"))
	orderHistoryConnection = config.GetCollection(config.DB, os.Getenv("ORDER_HISTORY"))
)

type OrderRepository interface {
	Init() (context.Context, context.CancelFunc)
	PlaceSingleOrder(order models.Orders) (*mongo.InsertOneResult, error)
	AddToCartOrder(order models.OrderCarts) error
	CheckDuplicateOrderToAddCart(orderId, userId primitive.ObjectID) error
	DeleteCartOrder(orderId primitive.ObjectID) error
	DeleteOrder(orderId primitive.ObjectID) error
	UpdateOrderQuantityAndPrice(order models.Orders) error

	UserCartItem(userId primitive.ObjectID) ([]models.OrderCarts, error)
	UpdateOrderStatus(status string, orderId primitive.ObjectID) error
	CreateOrderTrack(orderTrack models.OrderTrack) error
	CreateOrderHistory(orderData models.OrderHistory) error

	GetAllOrderTrack(orderId primitive.ObjectID) ([]models.OrderTrack, error)
	GetOrderById(orderId primitive.ObjectID) (models.Orders, error)
	DeleteOrderFromTrack(orderId []primitive.ObjectID) error
}

type orderRepository struct {
	ordersCollection       *mongo.Collection
	orderCartCollection    *mongo.Collection
	orderTrackCollection   *mongo.Collection
	orderHistoryCollection *mongo.Collection
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		ordersCollection:       ordersConnection,
		orderCartCollection:    orderCartConnection,
		orderTrackCollection:   orderTrackConnection,
		orderHistoryCollection: orderHistoryConnection,
	}
}

package repositories

import (
	"context"

	"github.com/aniket0951/order-services/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type OrderRepository interface {
	Init() (context.Context, context.CancelFunc)
	PlaceSingleOrder(order models.Orders) 
	AddToCartOrder(order models.OrderCarts) 
	CheckDuplicateOrderToAddCart(orderId, userId primitive.ObjectID) error
	DeleteCartOrder(orderId primitive.ObjectID) error
	DeleteOrder(orderId primitive.ObjectID) error
	UpdateOrderQuantityAndPrice(order models.Orders) error

	UserCartItem(userId primitive.ObjectID) ([]models.OrderCarts, error)
	UpdateOrderStatus(status string, orderId primitive.ObjectID) error
	CreateOrderTrack(orderTrack models.OrderTrack) 
	CreateOrderHistory(orderData models.OrderHistory) 

	GetAllOrderTrack(orderId primitive.ObjectID) ([]models.OrderTrack, error)
	GetOrderById(orderId primitive.ObjectID) (models.Orders, error)
	DeleteOrderFromTrack(orderId []primitive.ObjectID) 
}

type orderRepository struct {
	Orders []models.Orders
	OrderCarts []models.OrderCarts
	OrderTrack []models.OrderTrack
	OrderHistory []models.OrderHistory
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		Orders: []models.Orders{},
		OrderCarts: []models.OrderCarts{},
		OrderTrack: []models.OrderTrack{},
		OrderHistory: []models.OrderHistory{},
	}
}

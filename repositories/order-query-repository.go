package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/aniket0951/order-services/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *orderRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

// place a single order
func (db *orderRepository) PlaceSingleOrder(order models.Orders)  {
	db.Orders = append(db.Orders, order)
}

func (db *orderRepository) AddToCartOrder(order models.OrderCarts) {
	db.OrderCarts = append(db.OrderCarts, order)
}

// validate order can not be in cart already
func (db *orderRepository) CheckDuplicateOrderToAddCart(orderId, userId primitive.ObjectID) error {

	for i := range db.OrderCarts {
		if db.OrderCarts[i].OrderId == orderId && db.OrderCarts[i].UserId == userId {
			return errors.New("order already in cart")
		}
	}

	return nil
}

// remove oder from cart
func (db *orderRepository) DeleteCartOrder(orderId primitive.ObjectID) error {

	for i := range db.OrderCarts {
		if db.OrderCarts[i].OrderId == orderId {
			db.OrderCarts = append(db.OrderCarts[:i], db.OrderCarts[i+1:]...)
			return nil
		}
	}

	return nil
}

// delete original order from main collection
func (db *orderRepository) DeleteOrder(orderId primitive.ObjectID) error {
	for i := range db.Orders {
		if db.Orders[i].Id == orderId {
			db.Orders = append(db.Orders[:i], db.Orders[i+1:]...)
			return nil
		}
	}

	return nil
}

// user can update order before placed like to decrease quantity or change color
func (db *orderRepository) UpdateOrderQuantityAndPrice(order models.Orders) error {
	orders := db.Orders
	for i := range orders {
		if orders[i].Id == order.Id{
			orders[i].Quantity = order.Quantity
			orders[i].TotalPrice = order.TotalPrice

			return nil
		}
	}

	return errors.New("order not found to update")
}

func (db *orderRepository) UserCartItem(userId primitive.ObjectID) ([]models.OrderCarts, error) {
	orderCart := db.OrderCarts
	res := []models.OrderCarts{}
	for i := range orderCart {
		if orderCart[i].OrderStatus == "CART" && orderCart[i].UserId == userId {
			res = append(res, orderCart[i])
		}
	}

	return res, nil
}

// update order status live in orders collection
func (db *orderRepository) UpdateOrderStatus(status string, orderId primitive.ObjectID) error {

	order := db.Orders

	for i := range order {
		if order[i].Id == orderId{
			order[i].OrderStatus = status
			return nil
		}
	}
	return errors.New("order not found for update status")
}

// create order track
func (db *orderRepository) CreateOrderTrack(orderTrack models.OrderTrack)  {
	db.OrderTrack = append(db.OrderTrack, orderTrack)
}

func (db *orderRepository) CreateOrderHistory(orderData models.OrderHistory)  {
	db.OrderHistory = append(db.OrderHistory, orderData)
}

func (db *orderRepository) GetAllOrderTrack(orderId primitive.ObjectID) ([]models.OrderTrack, error) {

	order := db.OrderTrack
	res := []models.OrderTrack{}
	for i := range order {
		if order[i].OrderId == orderId {
			res = append(res, order[i])
		}
	}

	return res, nil
}

func (db *orderRepository) GetOrderById(orderId primitive.ObjectID) (models.Orders, error) {

	order := db.Orders

	for i := range order {
		if order[i].Id == orderId {
			return order[i], nil
		}
	}

	return models.Orders{}, nil
}

func (db *orderRepository) DeleteOrderFromTrack(orderId []primitive.ObjectID)  {

	order := db.OrderTrack

	for i := range order {
		for j := range orderId {
			if order[i].Id == orderId[j] {
				order = append(order[:i], order[i+1:]...)
			}
		}

	}
}

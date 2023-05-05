package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aniket0951/order-services/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *orderRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

// place a single order
func (db *orderRepository) PlaceSingleOrder(order models.Orders) (*mongo.InsertOneResult, error) {
	ctx, cancel := db.Init()
	defer cancel()

	res, err := db.ordersCollection.InsertOne(ctx, order)
	fmt.Println("inserted id from repo : ", res.InsertedID)
	return res, err
}

func (db *orderRepository) AddToCartOrder(order models.OrderCarts) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.orderCartCollection.InsertOne(ctx, order)
	return err
}

// validate order can not be in cart already
func (db *orderRepository) CheckDuplicateOrderToAddCart(orderId, userId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "order_id", Value: orderId},
		bson.E{Key: "user_id", Value: userId},
	}

	var order models.OrderCarts

	db.orderCartCollection.FindOne(ctx, filter).Decode(&order)

	if (order == models.OrderCarts{}) {
		return nil
	}
	return errors.New("order already in cart")
}

// remove oder from cart
func (db *orderRepository) DeleteCartOrder(orderId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "order_id", Value: orderId},
	}

	_, err := db.orderCartCollection.DeleteOne(ctx, filter)
	return err
}

// delete original order from main collection
func (db *orderRepository) DeleteOrder(orderId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: orderId},
	}
	_, err := db.ordersCollection.DeleteOne(ctx, filter)
	return err
}

// user can update order before placed like to decrease quantity or change color
func (db *orderRepository) UpdateOrderQuantityAndPrice(order models.Orders) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: order.Id},
	}

	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "quantity", Value: order.Quantity},
			bson.E{Key: "total_price", Value: order.TotalPrice},
		}},
	}

	res, err := db.ordersCollection.UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return errors.New("order not found to update")
	}
	return err
}

func (db *orderRepository) UserCartItem(userId primitive.ObjectID) ([]models.OrderCarts, error) {
	filter := bson.D{
		bson.E{Key: "order_status", Value: "CART"},
		bson.E{Key: "user_id", Value: userId},
	}

	ctx, cancel := db.Init()
	defer cancel()

	cursor, curErr := db.orderCartCollection.Find(ctx, filter)

	if curErr != nil {
		return nil, curErr
	}

	var orderCarts []models.OrderCarts

	if err := cursor.All(context.TODO(), &orderCarts); err != nil {
		return nil, err
	}

	return orderCarts, nil
}

// update order status live in orders collection
func (db *orderRepository) UpdateOrderStatus(status string, orderId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: orderId},
	}

	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "order_status", Value: status},
		}},
	}

	_, err := db.ordersCollection.UpdateOne(ctx, filter, update)
	return err
}

// create order track
func (db *orderRepository) CreateOrderTrack(orderTrack models.OrderTrack) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.orderTrackCollection.InsertOne(ctx, orderTrack)
	return err
}

func (db *orderRepository) CreateOrderHistory(orderData models.OrderHistory) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.orderHistoryCollection.InsertOne(ctx, orderData)
	return err
}

func (db *orderRepository) GetAllOrderTrack(orderId primitive.ObjectID) ([]models.OrderTrack, error) {
	filter := bson.D{
		bson.E{Key: "order_id", Value: orderId},
	}

	ctx, cancel := db.Init()
	defer cancel()

	cursor, curErr := db.orderTrackCollection.Find(ctx, filter)

	if curErr != nil {
		return nil, curErr
	}

	var trackData []models.OrderTrack

	if err := cursor.All(context.TODO(), &trackData); err != nil {
		return nil, err
	}

	return trackData, nil
}

func (db *orderRepository) GetOrderById(orderId primitive.ObjectID) (models.Orders, error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: orderId},
	}

	ctx, cancel := db.Init()
	defer cancel()

	var order models.Orders

	err := db.ordersCollection.FindOne(ctx, filter).Decode(&order)

	return order, err
}

func (db *orderRepository) DeleteOrderFromTrack(orderId []primitive.ObjectID) error {
	filter := bson.D{
		bson.E{Key: "order_id", Value: bson.D{
			bson.E{Key: "$in", Value: orderId},
		}},
	}

	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.orderTrackCollection.DeleteMany(ctx, filter)
	return err
}

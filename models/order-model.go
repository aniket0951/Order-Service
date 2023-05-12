package models

import (
	"time"

	"github.com/aniket0951/order-services/dto"
	"github.com/aniket0951/order-services/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductSellingID primitive.ObjectID `json:"prod_selling_id" bson:"prod_selling_id"`
	ProductId        primitive.ObjectID `json:"prod_id" bson:"prod_id"`
	Quantity         int64              `json:"quantity" bson:"quantity"`
	Category         string             `json:"category" bson:"category"`
	TotalPrice       float64            `json:"total_price" bson:"total_price"`
	UserId           primitive.ObjectID `json:"user_id" bson:"user_id"`
	OrderStatus      string             `json:"order_status" bson:"order_status"`
	CreatedAt        primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt        primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func (order *Orders) SetPlaceOrder(moduleData dto.CreateOrderDTO, tag string) Orders {
	newOrder := Orders{}

	prodSellId, _ := helper.ConvertStringToPrimitive(moduleData.ProductSellingID)
	userId, _ := helper.ConvertStringToPrimitive(moduleData.UserId)
	prodId, _ := helper.ConvertStringToPrimitive(moduleData.ProductId)

	newOrder.Id = primitive.NewObjectID()
	newOrder.ProductSellingID = prodSellId
	newOrder.ProductId = prodId
	newOrder.UserId = userId
	newOrder.Quantity = moduleData.Quantity
	newOrder.Category = moduleData.Category
	newOrder.OrderStatus = tag
	newOrder.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newOrder.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newOrder
}

type OrderCarts struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrderId   primitive.ObjectID `json:"order_id" bson:"order_id"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
	OrderStatus string `json:"order_status" bson:"order_status"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type OrderTrack struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrderId     primitive.ObjectID `json:"order_id" bson:"order_id"`
	OrderStatus string             `json:"order_status" bson:"order_status"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func (track *OrderTrack) SetOrderTrack(orderID primitive.ObjectID, status string) OrderTrack {
	newOrderTrack := OrderTrack{}

	newOrderTrack.OrderId = orderID
	newOrderTrack.OrderStatus = status
	newOrderTrack.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newOrderTrack.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newOrderTrack
}

type OrderHistory struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Order      Orders             `json:"order" bson:"order"`
	OrderTrack []OrderTrack       `json:"order_track" bson:"order_track"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func (trackHistory *OrderHistory) SetOrderHistory(order Orders, orderTrack []OrderTrack) OrderHistory {
	newOrderHistory := OrderHistory{}

	newOrderHistory.Order = order
	newOrderHistory.OrderTrack = orderTrack
	newOrderHistory.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newOrderHistory.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newOrderHistory
}

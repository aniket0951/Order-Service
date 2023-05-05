package services

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/aniket0951/order-services/dto"
	"github.com/aniket0951/order-services/helper"
	"github.com/aniket0951/order-services/models"
	"github.com/aniket0951/order-services/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	PlaceSingleOrder(order dto.CreateOrderDTO) error
	AddToCart(order dto.CreateOrderDTO) error
	RemoveItemFromCart(orderId string) error
	UpdateOrderStatus(orderId, status string) error
	CancelOrder(orderId string) error
}

type orderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

// placing single product order
func (ser *orderService) PlaceSingleOrder(order dto.CreateOrderDTO) error {
	if order.Quantity > 10 {
		return errors.New("quantity should be less than 10")
	}

	_, valErr := helper.ValidatePrimitiveId(order.ProductId)
	_, valSellIdErr := helper.ValidatePrimitiveId(order.ProductSellingID)

	if valErr != nil {
		return valErr
	}

	if valSellIdErr != nil {
		return valSellIdErr
	}

	orderToPlace := new(models.Orders).SetPlaceOrder(order, helper.PLACED)
	orderToPlace.TotalPrice = float64(order.Price) * float64(order.Quantity)
	_, err := ser.orderRepo.PlaceSingleOrder(orderToPlace)
	if err != nil {
		return err
	}

	// updating product count by quantity
	quantityStr := strconv.Itoa(int(order.Quantity))
	countErr := ser.UpdateProductCount("decrease", quantityStr, order.ProductId)

	return countErr
}

func (ser *orderService) AddToCart(order dto.CreateOrderDTO) error {
	if order.Quantity > 10 {
		return errors.New("quantity should be less than 10")
	}

	_, valErr := helper.ValidatePrimitiveId(order.ProductId)
	_, valSellIdErr := helper.ValidatePrimitiveId(order.ProductSellingID)
	userID, valUserIdErr := helper.ValidatePrimitiveId(order.UserId)

	if valErr != nil {
		return valErr
	}

	if valSellIdErr != nil {
		return valSellIdErr
	}

	if valUserIdErr != nil {
		return valUserIdErr
	}

	orderToPlace := new(models.Orders).SetPlaceOrder(order, helper.CART)
	orderToPlace.TotalPrice = float64(order.Price) * float64(order.Quantity)
	_, err := ser.orderRepo.PlaceSingleOrder(orderToPlace)
	if err != nil {
		return err
	}

	orderItem := models.OrderCarts{
		OrderId:   orderToPlace.Id,
		UserId:    userID,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	return ser.orderRepo.AddToCartOrder(orderItem)
}

// update product counts accordingly
func (ser *orderService) UpdateProductCount(tag string, num string, productId string) error {
	reqUrl := helper.INCREASE_DECREASE_PRODUCT + "?tag=" + tag + "&number=" + num + "&product_id=" + productId

	req, err := http.NewRequest(http.MethodPut, reqUrl, &bytes.Buffer{})

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(body)
	return nil
}

func (ser *orderService) RemoveItemFromCart(orderId string) error {
	orderObjId, err := helper.ValidatePrimitiveId(orderId)

	if err != nil {
		return err
	}

	// remove item from cart
	rmErr := ser.orderRepo.DeleteCartOrder(orderObjId)
	if rmErr != nil {
		return rmErr
	}

	// after remove the cart item remove item data from orders
	return ser.orderRepo.DeleteOrder(orderObjId)
}

func (ser *orderService) CreateOrderTrack(order models.Orders) error {
	orderTrack := models.OrderTrack{}
	orderTrack.OrderId = order.Id
	orderTrack.OrderStatus = order.OrderStatus
	orderTrack.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	orderTrack.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return ser.orderRepo.CreateOrderTrack(orderTrack)
}

func (ser *orderService) UpdateOrderStatus(orderId, status string) error {
	orderObjId, valErr := helper.ValidatePrimitiveId(orderId)
	if valErr != nil {
		return valErr
	}

	err := ser.orderRepo.UpdateOrderStatus(status, orderObjId)

	if err != nil {
		return err
	}

	orderTrack := new(models.OrderTrack).SetOrderTrack(orderObjId, status)

	trackErr := ser.orderRepo.CreateOrderTrack(orderTrack)

	if status == helper.COMPLETED {
		ser.CreateOrderHistory(orderObjId)
	}

	return trackErr
}

// make a order history after order has been completed
func (ser *orderService) CreateOrderHistory(orderId primitive.ObjectID) error {
	order, orderErr := ser.orderRepo.GetOrderById(orderId)

	if orderErr != nil {
		return orderErr
	}

	orderTrack, trackErr := ser.orderRepo.GetAllOrderTrack(orderId)

	if trackErr != nil {
		return trackErr
	}

	orderHistory := new(models.OrderHistory).SetOrderHistory(order, orderTrack)

	// create a track history
	hisErr := ser.orderRepo.CreateOrderHistory(orderHistory)

	if hisErr != nil {
		return hisErr
	}

	// remove all data
	delErr := ser.orderRepo.DeleteOrder(orderId)

	if delErr != nil {
		return delErr
	}

	delTrackErr := ser.orderRepo.DeleteOrderFromTrack([]primitive.ObjectID{orderId})
	return delTrackErr
}

// if order get cancel then increase a product count
func (ser *orderService) CancelOrder(orderId string) error {
	orderObjId, valErr := helper.ValidatePrimitiveId(orderId)

	if valErr != nil {
		return valErr
	}

	order, orderErr := ser.orderRepo.GetOrderById(orderObjId)

	if orderErr != nil {
		return orderErr
	}

	upErr := ser.UpdateOrderStatus(orderId, helper.CANCELLED)
	if upErr != nil {
		return upErr
	}

	quantity := strconv.Itoa(int(order.Quantity))

	err := ser.UpdateProductCount("increase", quantity, order.ProductId.String())
	return err
}

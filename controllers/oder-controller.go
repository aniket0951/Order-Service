package controllers

import (
	"errors"
	"net/http"

	"github.com/aniket0951/order-services/dto"
	"github.com/aniket0951/order-services/helper"
	"github.com/aniket0951/order-services/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderControllers interface {
	PlaceSingleOrder(*gin.Context)
	AddToCart(*gin.Context)
	RemoveItemFromCart(*gin.Context)
	UpdateOrderStatus(*gin.Context)
	CancelOrder(*gin.Context)
}

type orderControllers struct {
	orderService services.OrderService
}

func NewOrderControllers(orderService services.OrderService) OrderControllers {
	return &orderControllers{
		orderService: orderService,
	}
}

func (c *orderControllers) PlaceSingleOrder(ctx *gin.Context) {
	order := dto.CreateOrderDTO{}
	_ = ctx.BindJSON(&order)

	if (order == dto.CreateOrderDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&order); helper.CheckError(stErr, ctx) {
		return
	}

	err := c.orderService.PlaceSingleOrder(order)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("order has been placed successfully", helper.EmptyObj{}, helper.ORDER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderControllers) AddToCart(ctx *gin.Context) {
	order := dto.CreateOrderDTO{}
	_ = ctx.BindJSON(&order)

	if (order == dto.CreateOrderDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&order); helper.CheckError(stErr, ctx) {
		return
	}

	err := c.orderService.AddToCart(order)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("order has been added into cart successfully", helper.EmptyObj{}, helper.ORDER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderControllers) RemoveItemFromCart(ctx *gin.Context) {
	orderId := helper.GetRequestQueryParam("order_id", ctx)

	if helper.CheckRequestParamEmpty(orderId, ctx) {
		return
	}

	err := c.orderService.RemoveItemFromCart(orderId)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("item remove success", helper.EmptyObj{}, helper.ORDER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderControllers) UpdateOrderStatus(ctx *gin.Context) {
	orderId := helper.GetRequestQueryParam("order_id", ctx)
	status := helper.GetRequestQueryParam("status", ctx)

	if helper.CheckRequestParamEmpty(orderId, ctx) || helper.CheckRequestParamEmpty(status, ctx) {
		return
	}

	if status != helper.DISPATCHED || status != helper.COMPLETED || status != helper.CART {
		helper.BuildUnProcessableEntity(ctx, errors.New("invalid status passed"))
		return
	}

	err := c.orderService.UpdateOrderStatus(orderId, status)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.UPDATE_SUCCESS, helper.EmptyObj{}, helper.ORDER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderControllers) CancelOrder(ctx *gin.Context) {
	orderId := helper.GetRequestQueryParam("order_id", ctx)

	if helper.CheckRequestParamEmpty(orderId, ctx) {
		return
	}

	err := c.orderService.CancelOrder(orderId)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("order has been cancel successfully", helper.EmptyObj{}, helper.ORDER_DATA)
	ctx.JSON(http.StatusOK, response)
}

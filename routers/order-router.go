package routers

import (
	"github.com/aniket0951/order-services/controllers"
	"github.com/aniket0951/order-services/repositories"
	"github.com/aniket0951/order-services/services"
	"github.com/gin-gonic/gin"
)

var (
	orderrepo       = repositories.NewOrderRepository()
	orderservice    = services.NewOrderService(orderrepo)
	ordercontroller = controllers.NewOrderControllers(orderservice)
)

func OrderRouter(router *gin.Engine) {
	orderRoutes := router.Group("/api/order")
	{
		orderRoutes.POST("/place-single-order", ordercontroller.PlaceSingleOrder)
		orderRoutes.POST("/add-to-card", ordercontroller.AddToCart)
		orderRoutes.DELETE("/remove-item-from-cart", ordercontroller.RemoveItemFromCart)
		orderRoutes.PUT("/update-order-status", ordercontroller.UpdateOrderStatus)
		orderRoutes.PUT("/cancel-order", ordercontroller.CancelOrder)
	}

}

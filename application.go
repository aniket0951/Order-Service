package main

import (
	"github.com/aniket0951/order-services/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.SetTrustedProxies(nil)


	router.Static("static", "static")

	routers.OrderRouter(router)

	router.Run(":8080")
}

package routes

import (
	// golang package
	"order-service/cmd/order/handler"
	"order-service/middleware"

	// external package
	"github.com/gin-gonic/gin"
)

// SetupRoutes setup routes by given router pointer of gin.Engine, and OrderHandler.
func SetupRoutes(router *gin.Engine, orderHandler handler.OrderHandler, jwtSecret string) {
	router.Use(middleware.RequestLogger())
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	router.Use(authMiddleware)
	router.POST("/v1/checkout", orderHandler.Checkout)

	router.GET("/v1/order_history", orderHandler.GetOrderHistory)
}

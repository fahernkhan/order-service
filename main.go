package main

import (
	// golang package
	"context"
	"order-service/cmd/order/handler"
	"order-service/cmd/order/repository"
	"order-service/cmd/order/resource"
	"order-service/cmd/order/service"
	"order-service/cmd/order/usecase"
	"order-service/config"
	"order-service/infrastructure/log"
	"order-service/kafka"
	"order-service/kafka/consumer"
	"order-service/routes"

	// external package
	"github.com/gin-gonic/gin"
)

// main main.
func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDB(&cfg)

	log.SetupLogger()
	kafkaProducer := kafka.NewKafkaProducer([]string{"localhost:9093"})
	defer kafkaProducer.Close()

	orderRepository := repository.NewOrderRepository(db, redis)
	orderService := service.NewOrderService(*orderRepository)
	orderUsecase := usecase.NewOrderUsecase(*orderService, *kafkaProducer)
	orderHandler := handler.NewOrderHandler(*orderUsecase)

	port := cfg.App.Port
	router := gin.Default()
	routes.SetupRoutes(router, *orderHandler, cfg.Secret.JWTSecret)
	router.Run(":" + port)

	// kafka consumer
	kafkaPaymentSuccessConsumer := consumer.NewPaymentSuccessConsumer(
		[]string{"localhost:9093"},
		"payment.success",
		*orderService,
		*kafkaProducer,
	)

	kafkaPaymentSuccessConsumer.StartPaymentSuccessConsumer(context.Background())

	kafkaPaymentFailedConsumer := consumer.NewPaymentFailedConsumer(
		[]string{"localhost:9093"},
		"payment.failed",
		*orderService,
		*kafkaProducer,
	)

	kafkaPaymentFailedConsumer.Start(context.Background())

	log.Logger.Printf("Server running on port: %s", port)
}

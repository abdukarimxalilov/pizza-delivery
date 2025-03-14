package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/abdukarimxalilov/order-pizza/config"
	"github.com/abdukarimxalilov/order-pizza/constants"
	"github.com/abdukarimxalilov/order-pizza/handler"
	"github.com/abdukarimxalilov/order-pizza/logger"
	"github.com/abdukarimxalilov/order-pizza/middleware"
	"github.com/abdukarimxalilov/order-pizza/routes"
	"github.com/abdukarimxalilov/order-pizza/service"
)

func main() {

	app := gin.Default()

	app.Use(gin.Recovery())

	app.Use(middleware.CorsMiddleware)

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":    "Pizza Shop is open",
			"statusCode": 200,
		})
	})

	messagePublisher := service.GetMessagePublisherService()
	messageConsumer := service.GetMessageConsumerService()

	websocketHandler := handler.GetNewWebSocketHandler()
	messageProcessor := service.GetMessageProcessorService(messagePublisher, websocketHandler.GetConnectionMap())

	go func() {
		err := messageConsumer.ConsumeEventAndProcess(constants.KITCHEN_ORDER_QUEUE, messageProcessor)
		if err != nil {
			logger.Log(fmt.Sprintf("failed to consume events : %v", err))
		}
	}()

	routes.RegisterRoutes(app, messagePublisher, websocketHandler)

	port := config.GetEnvProperty("port")
	logger.Log(fmt.Sprintf("Pizza shop started successfully on port : %s", port))

	app.Run(fmt.Sprintf(":%s", port))

}
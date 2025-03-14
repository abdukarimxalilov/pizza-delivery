package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/abdukarimxalilov/order-pizza/handler"
	"github.com/abdukarimxalilov/order-pizza/service"
)

func RegisterOrderRoutes(router *gin.RouterGroup, messagePublisher service.IMessagePublisher) {

	oh := handler.GetOrderHandler(messagePublisher)

	router.POST(
		"/create",
		oh.CreateOrder,
	)
}
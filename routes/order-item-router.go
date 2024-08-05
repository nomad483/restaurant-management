package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

func OrderItemRoutes(incomingRoute *gin.Engine) {
	orderItemGroup := incomingRoute.Group("/api/order-items")

	orderItemGroup.GET("/", controllers.GetOrderItems())
	orderItemGroup.GET("/:id", controllers.GetOrderItem())
	orderItemGroup.GET("/order/:id", controllers.GetOrderItemByOrder())
	orderItemGroup.POST("/", controllers.CreateOrderItem())
	orderItemGroup.PATCH("/:id", controllers.UpdateOrderItem())
	orderItemGroup.DELETE("/:id", controllers.DeleteOrderItem())
}

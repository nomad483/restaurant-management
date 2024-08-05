package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

func OrderRoutes(router *gin.Engine) {
	orderGroup := router.Group("/api/orders")

	orderGroup.GET("/", controllers.GetOrders())
	orderGroup.GET("/:id", controllers.GetOrder())
	orderGroup.POST("/", controllers.CreateOrder())
	orderGroup.PATCH("/:id", controllers.UpdateOrder())
	orderGroup.DELETE("/:id", controllers.DeleteOrder())
}

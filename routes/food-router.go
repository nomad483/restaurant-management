package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

var foodController controllers.FoodController = controllers.NewFoodController()

func FoodRoutes(incomingRoutes *gin.Engine) {
	foodGroup := incomingRoutes.Group("/api/foods")

	foodGroup.GET("/", foodController.GetFoods())
	foodGroup.GET("/:id", foodController.GetFood())
	foodGroup.POST("/", foodController.CreateFood())
	foodGroup.PATCH("/:id", foodController.UpdateFood())
	foodGroup.DELETE("/:id", foodController.DeleteFood())
}

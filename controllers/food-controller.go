package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/models"
	"github.com/nomad483/restaurant-managment/services"
	"net/http"
)

var foodService services.FoodService = services.NewFoodService()

type FoodController interface {
	GetFoods() gin.HandlerFunc
	GetFood() gin.HandlerFunc
	CreateFood() gin.HandlerFunc
	UpdateFood() gin.HandlerFunc
	DeleteFood() gin.HandlerFunc
}

type foodController struct{}

func NewFoodController() FoodController {
	return &foodController{}
}

func (controller *foodController) GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseFoods, err := foodService.GetFoods()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, responseFoods)
	}
}

func (controller *foodController) GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		responseFood, err := foodService.GetFoodById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, responseFood)
	}
}

func (controller *foodController) CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		responseFood, err := foodService.CreateFood(food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responseFood)
	}
}

func (controller *foodController) UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var food models.UpdateFood
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedFood, err := foodService.UpdateFood(id, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedFood)
	}
}

func (controller *foodController) DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := foodService.DeleteFood(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food deleted"})
	}
}

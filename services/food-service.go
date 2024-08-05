package services

import (
	"github.com/nomad483/restaurant-managment/models"
	"github.com/nomad483/restaurant-managment/repositories"
)

var foodRepository repositories.FoodRepository = repositories.NewFoodRepository()

type FoodService interface {
	GetFoods() (models.ResponseFoodList, error)
	GetFoodById(id string) (models.ResponseFood, error)
	CreateFood(food models.Food) (models.ResponseFood, error)
	UpdateFood(id string, food models.UpdateFood) (models.ResponseFood, error)
	DeleteFood(id string) error

	mapFoodToResponseFood(food models.Food) models.ResponseFood
}

type foodService struct{}

func NewFoodService() FoodService {
	return &foodService{}
}

func (service *foodService) mapFoodToResponseFood(food models.Food) models.ResponseFood {
	return models.ResponseFood{
		ID:        food.ID,
		Name:      food.Name,
		Price:     food.Price,
		FoodImage: food.FoodImage,
		MenuId:    food.MenuId,
	}
}

func (service *foodService) GetFoods() (models.ResponseFoodList, error) {
	foods, err := foodRepository.GetFoods()
	if err != nil {
		return models.ResponseFoodList{}, err
	}

	var responseFoods []models.ResponseFood

	for _, food := range foods {
		responseFoods = append(responseFoods, service.mapFoodToResponseFood(food))
	}

	return models.ResponseFoodList{Data: responseFoods}, nil
}

func (service *foodService) GetFoodById(id string) (models.ResponseFood, error) {
	food, err := foodRepository.GetFoodById(id)
	if err != nil {
		return models.ResponseFood{}, err
	}

	return service.mapFoodToResponseFood(food), nil
}

func (service *foodService) CreateFood(food models.Food) (models.ResponseFood, error) {
	newFood, err := foodRepository.CreateFood(food)
	if err != nil {
		return models.ResponseFood{}, err
	}

	return service.mapFoodToResponseFood(newFood), nil
}

func (service *foodService) UpdateFood(id string, food models.UpdateFood) (models.ResponseFood, error) {
	updatedFood, err := foodRepository.UpdateFood(id, food)
	if err != nil {
		return models.ResponseFood{}, err
	}

	return service.mapFoodToResponseFood(updatedFood), nil
}

func (service *foodService) DeleteFood(id string) error {
	return foodRepository.DeleteFood(id)
}

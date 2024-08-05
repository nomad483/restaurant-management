package repositories

import (
	"context"
	"github.com/nomad483/restaurant-managment/database"
	"github.com/nomad483/restaurant-managment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

type FoodRepository interface {
	GetFoods() ([]models.Food, error)
	GetFoodById(id string) (models.Food, error)
	CreateFood(food models.Food) (models.Food, error)
	UpdateFood(id string, food models.UpdateFood) (models.Food, error)
	DeleteFood(id string) error
}

type foodRepository struct {
}

func NewFoodRepository() FoodRepository {
	return &foodRepository{}
}

func (repository *foodRepository) GetFoods() ([]models.Food, error) {
	var foods []models.Food
	result, err := foodCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if err = result.All(context.TODO(), &foods); err != nil {
		log.Fatal(err)
	}

	return foods, nil
}

func (repository *foodRepository) GetFoodById(id string) (models.Food, error) {
	var food models.Food
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return food, err
	}
	err = foodCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&food)
	if err != nil {
		return food, err
	}

	return food, nil
}

func (repository *foodRepository) CreateFood(food models.Food) (models.Food, error) {
	food.ID = primitive.NewObjectID()
	food.CreatedAt = time.Now()
	food.UpdatedAt = time.Now()

	_, err := foodCollection.InsertOne(context.TODO(), food)
	if err != nil {
		return models.Food{}, err
	}

	return food, nil
}

func (repository *foodRepository) UpdateFood(id string, food models.UpdateFood) (models.Food, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Food{}, err
	}

	updateData := bson.M{
		"updated_at": time.Now(),
	}

	if food.FoodImage != "" {
		updateData["food_image"] = food.FoodImage
	}

	if food.Name != "" {
		updateData["name"] = food.Name
	}

	if food.Price != 0 {
		updateData["price"] = food.Price
	}

	_, err = foodCollection.UpdateByID(context.TODO(), objectId, bson.M{"$set": updateData})

	if err != nil {
		return models.Food{}, err
	}

	updatedFood, err := repository.GetFoodById(id)

	return updatedFood, err
}

func (repository *foodRepository) DeleteFood(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = foodCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	return err
}

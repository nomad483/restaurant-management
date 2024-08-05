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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

type MenuRepository interface {
	GetMenus() ([]models.Menu, error)
	GetMenuBuId(id string) (models.Menu, error)
	CreateMenu(menu models.Menu) (models.Menu, error)
	UpdateMenu(id string, menu models.UpdateMenu) (models.Menu, error)
	DeleteMenu(id string) error
}

type menuRepository struct{}

func NewMenuRepository() MenuRepository {
	return &menuRepository{}
}

func (repository *menuRepository) GetMenus() ([]models.Menu, error) {
	var menus []models.Menu
	result, err := menuCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if err = result.All(context.TODO(), &menus); err != nil {
		log.Fatal(err)
	}

	return menus, nil
}

func (repository *menuRepository) GetMenuBuId(id string) (models.Menu, error) {
	var menu models.Menu
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return menu, err
	}

	err = menuCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&menu)
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (repository *menuRepository) CreateMenu(menu models.Menu) (models.Menu, error) {
	menu.ID = primitive.NewObjectID()
	menu.CreatedAt = time.Now()
	menu.UpdatedAt = time.Now()

	_, err := menuCollection.InsertOne(context.TODO(), menu)
	if err != nil {
		return models.Menu{}, err
	}

	return menu, nil
}

func (repository *menuRepository) UpdateMenu(id string, menu models.UpdateMenu) (models.Menu, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Menu{}, err
	}

	// Create a map for the fields to update
	updateData := bson.M{
		"updated_at": time.Now(),
	}

	// Add fields to update only if they are not empty or zero
	if menu.Name != "" {
		updateData["name"] = menu.Name
	}
	if menu.Category != "" {
		updateData["category"] = menu.Category
	}
	if menu.StartDate != nil {
		updateData["start_date"] = menu.StartDate
	}
	if menu.EndDate != nil {
		updateData["end_date"] = menu.EndDate
	}

	_, err = menuCollection.UpdateByID(context.TODO(), objectID, bson.M{"$set": updateData})

	if err != nil {
		return models.Menu{}, err
	}

	updatedMenu, err := repository.GetMenuBuId(id)

	return updatedMenu, err
}

func (repository *menuRepository) DeleteMenu(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = menuCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

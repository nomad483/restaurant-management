package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      *string            `json:"name" bson:"name" validate:"required,min=2,max=100" binding:"required,min=2,max=100"`
	Price     *float64           `json:"price" bson:"price" validate:"required" binding:"required"`
	FoodImage *string            `json:"food_image" bson:"food_image" validate:"required,url" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	FoodId    string             `json:"food_id" bson:"food_id"`
	MenuId    *string            `json:"menu_id" bson:"menu_id" validate:"required" binding:"required"`
}

type UpdateFood struct {
	Name      string  `json:"name" validate:"required,min=2,max=100" binding:"required,min=2,max=100"`
	Price     float64 `json:"price" validate:"required" binding:"required"`
	FoodImage string  `json:"food_image" validate:"required,url" binding:"required"`
}

type ResponseFood struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      *string            `json:"name"`
	Price     *float64           `json:"price"`
	FoodImage *string            `json:"food_image"`
	MenuId    *string            `json:"menu_id"`
}

type ResponseFoodList struct {
	Data []ResponseFood `json:"data"`
}

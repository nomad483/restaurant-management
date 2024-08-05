package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name" validate:"required" binding:"required"`
	Category  string             `json:"category" bson:"category" validate:"required" binding:"required"`
	StartDate *time.Time         `json:"start_date" bson:"start_date"`
	EndDate   *time.Time         `json:"end_date" bson:"end_date"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	MenuId    string             `json:"menu_id" bson:"menu_id"`
}

type UpdateMenu struct {
	Name      string     `json:"name" bson:"name"`
	Category  string     `json:"category" bson:"category"`
	StartDate *time.Time `json:"start_date" bson:"start_date"`
	EndDate   *time.Time `json:"end_date" bson:"end_date"`
}

type ResponseMenu struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Category  string             `json:"category"`
	StartDate *time.Time         `json:"start_date"`
	EndDate   *time.Time         `json:"end_date"`
}

type ResponseMenuList struct {
	Data []ResponseMenu `json:"data"`
}

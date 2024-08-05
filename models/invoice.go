package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	InvoiceId      string             `json:"invoice_id" bson:"invoice_id"`
	OrderId        string             `json:"order_id" bson:"order_id"`
	PaymentMethod  *string            `json:"payment_method" bson:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  *string            `json:"payment_status" bson:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate *time.Time         `json:"payment_due_date" bson:"payment_due_date"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}

type UpdateInvoice struct {
	OrderId        string     `json:"order_id" bson:"order_id"`
	PaymentMethod  *string    `json:"payment_method" bson:"payment_method"`
	PaymentStatus  *string    `json:"payment_status" bson:"payment_status"`
	PaymentDueDate *time.Time `json:"payment_due_date" bson:"payment_due_date"`
}

type ResponseInvoice struct {
	ID             primitive.ObjectID `json:"_id"`
	OrderId        string             `json:"order_id"`
	PaymentMethod  *string            `json:"payment_method"`
	PaymentStatus  *string            `json:"payment_status"`
	PaymentDueDate *time.Time         `json:"payment_due_date"`
}

type ResponseInvoiceList struct {
	Data []ResponseInvoice `json:"data"`
}

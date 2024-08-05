package repositories

import (
	"context"
	"github.com/nomad483/restaurant-managment/database"
	"github.com/nomad483/restaurant-managment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

type InvoiceRepository interface {
	GetInvoices() ([]models.Invoice, error)
	GetInvoiceByID(id string) (models.Invoice, error)
	CreateInvoice(invoice models.Invoice) (models.Invoice, error)
	UpdateInvoice(id string, invoice models.UpdateInvoice) (models.Invoice, error)
	DeleteInvoice(id string) error
}

type invoiceRepository struct{}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{}
}

func (repository *invoiceRepository) GetInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	cursor, err := invoiceCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &invoices); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (repository *invoiceRepository) GetInvoiceByID(id string) (models.Invoice, error) {
	var invoice models.Invoice
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return invoice, err
	}

	err = invoiceCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&invoice)
	if err != nil {
		return invoice, err
	}

	return invoice, nil
}

func (repository *invoiceRepository) CreateInvoice(invoice models.Invoice) (models.Invoice, error) {
	invoice.ID = primitive.NewObjectID()
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()

	_, err := invoiceCollection.InsertOne(context.TODO(), invoice)
	if err != nil {
		return models.Invoice{}, err
	}

	return invoice, nil
}

func (repository *invoiceRepository) UpdateInvoice(id string, invoice models.UpdateInvoice) (models.Invoice, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Invoice{}, err
	}

	updatedData := bson.M{
		"updated_at": time.Now(),
	}

	if invoice.OrderId != "" {
		updatedData["order_id"] = invoice.OrderId
	}

	if invoice.PaymentMethod != nil {
		updatedData["payment_method"] = invoice.PaymentMethod
	}

	if invoice.PaymentStatus != nil {
		updatedData["payment_status"] = invoice.PaymentStatus
	}

	if invoice.PaymentDueDate != nil {
		updatedData["payment_due_date"] = invoice.PaymentDueDate
	}

	_, err = invoiceCollection.UpdateByID(context.TODO(), objectID, bson.M{"$set": updatedData})
	if err != nil {
		return models.Invoice{}, err
	}

	updatedInvoice, err := repository.GetInvoiceByID(id)

	return updatedInvoice, err
}

func (repository *invoiceRepository) DeleteInvoice(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = invoiceCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

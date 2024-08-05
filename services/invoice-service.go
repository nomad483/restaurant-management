package services

import (
	"github.com/nomad483/restaurant-managment/models"
	"github.com/nomad483/restaurant-managment/repositories"
)

var invoiceRepository repositories.InvoiceRepository = repositories.NewInvoiceRepository()

type InvoiceService interface {
	GetInvoices() (models.ResponseInvoiceList, error)
	GetInvoiceById(id string) (models.ResponseInvoice, error)
	CreateInvoice(invoice models.Invoice) (models.ResponseInvoice, error)
	UpdateInvoice(id string, invoice models.UpdateInvoice) (models.ResponseInvoice, error)
	DeleteInvoice(id string) error

	mapInvoiceToResponseInvoice(invoice models.Invoice) models.ResponseInvoice
}

type invoiceService struct{}

func NewInvoiceService() InvoiceService {
	return &invoiceService{}
}

func (service *invoiceService) mapInvoiceToResponseInvoice(invoice models.Invoice) models.ResponseInvoice {
	return models.ResponseInvoice{
		ID:             invoice.ID,
		OrderId:        invoice.OrderId,
		PaymentMethod:  invoice.PaymentMethod,
		PaymentStatus:  invoice.PaymentStatus,
		PaymentDueDate: invoice.PaymentDueDate,
	}
}

func (service *invoiceService) GetInvoices() (models.ResponseInvoiceList, error) {
	invoices, err := invoiceRepository.GetInvoices()
	if err != nil {
		return models.ResponseInvoiceList{}, err
	}

	var responseInvoices []models.ResponseInvoice

	for _, invoice := range invoices {
		responseInvoices = append(responseInvoices, service.mapInvoiceToResponseInvoice(invoice))
	}

	return models.ResponseInvoiceList{Data: responseInvoices}, nil
}

func (service *invoiceService) GetInvoiceById(id string) (models.ResponseInvoice, error) {
	invoice, err := invoiceRepository.GetInvoiceByID(id)
	if err != nil {
		return models.ResponseInvoice{}, err
	}

	return service.mapInvoiceToResponseInvoice(invoice), nil
}

func (service *invoiceService) CreateInvoice(invoice models.Invoice) (models.ResponseInvoice, error) {
	newInvoice, err := invoiceRepository.CreateInvoice(invoice)
	if err != nil {
		return models.ResponseInvoice{}, err
	}

	return service.mapInvoiceToResponseInvoice(newInvoice), nil
}

func (service *invoiceService) UpdateInvoice(id string, invoice models.UpdateInvoice) (models.ResponseInvoice, error) {
	updatedInvoice, err := invoiceRepository.UpdateInvoice(id, invoice)
	if err != nil {
		return models.ResponseInvoice{}, err
	}

	return service.mapInvoiceToResponseInvoice(updatedInvoice), nil
}

func (service *invoiceService) DeleteInvoice(id string) error {
	return invoiceRepository.DeleteInvoice(id)
}

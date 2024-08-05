package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/models"
	"github.com/nomad483/restaurant-managment/services"
	"net/http"
)

var invoiceService services.InvoiceService = services.NewInvoiceService()

type InvoiceController interface {
	GetInvoices() gin.HandlerFunc
	GetInvoice() gin.HandlerFunc
	CreateInvoice() gin.HandlerFunc
	UpdateInvoice() gin.HandlerFunc
	DeleteInvoice() gin.HandlerFunc
}

type invoiceController struct{}

func NewInvoiceController() InvoiceController {
	return &invoiceController{}
}

func (controller *invoiceController) GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseInvoices, err := invoiceService.GetInvoices()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responseInvoices)
	}
}

func (controller *invoiceController) GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		responceInvoice, err := invoiceService.GetInvoiceById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responceInvoice)
	}
}

func (controller *invoiceController) CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice
		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		responseInvoice, err := invoiceService.CreateInvoice(invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responseInvoice)
	}
}

func (controller *invoiceController) UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var invoice models.UpdateInvoice
		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedInvoice, err := invoiceService.UpdateInvoice(id, invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedInvoice)
	}
}

func (controller *invoiceController) DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := invoiceService.DeleteInvoice(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
	}
}

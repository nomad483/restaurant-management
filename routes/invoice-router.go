package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

var invoiceController controllers.InvoiceController = controllers.NewInvoiceController()

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	invoiceGroup := incomingRoutes.Group("/api/invoices")

	invoiceGroup.GET("/", invoiceController.GetInvoices())
	invoiceGroup.GET("/:id", invoiceController.GetInvoice())
	invoiceGroup.POST("/", invoiceController.CreateInvoice())
	invoiceGroup.PATCH("/:id", invoiceController.UpdateInvoice())
	invoiceGroup.DELETE("/:id", invoiceController.DeleteInvoice())
}

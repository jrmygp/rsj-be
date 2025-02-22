package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/shipment"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShipmentController struct {
	service services.Service
}

func NewShipmentController(service services.Service) *ShipmentController {
	return &ShipmentController{service}
}

func convertShipmentResponse(o models.Shipment) responses.ShipmentResponse {
	var quotationResponse []responses.LinkedDocResponse
	var exportResponse []responses.LinkedDocResponse
	var importResponse []responses.LinkedDocResponse
	var d2dResponse []responses.LinkedDocResponse

	for _, quotation := range o.Quotations {
		itemResponse := responses.LinkedDocResponse{
			ID:        quotation.ID,
			DocNumber: quotation.QuotationNumber,
		}
		quotationResponse = append(quotationResponse, itemResponse)
	}

	for _, export := range o.InvoiceExports {
		itemResponse := responses.LinkedDocResponse{
			ID:        export.ID,
			DocNumber: export.InvoiceNumber,
		}
		exportResponse = append(exportResponse, itemResponse)
	}

	for _, item := range o.InvoiceImports {
		itemResponse := responses.LinkedDocResponse{
			ID:        item.ID,
			DocNumber: item.InvoiceNumber,
		}
		importResponse = append(importResponse, itemResponse)
	}

	for _, d2d := range o.InvoiceDoorToDoors {
		itemResponse := responses.LinkedDocResponse{
			ID:        d2d.ID,
			DocNumber: d2d.InvoiceNumber,
		}
		d2dResponse = append(d2dResponse, itemResponse)
	}

	return responses.ShipmentResponse{
		ID:                 o.ID,
		ShipmentNumber:     o.ShipmentNumber,
		Quotations:         quotationResponse,
		InvoiceExports:     exportResponse,
		InvoiceImports:     importResponse,
		InvoiceDoorToDoors: d2dResponse,
	}
}

func (h *ShipmentController) FindAll(c *gin.Context) {
	// Search
	searchQuery := c.Query("search")

	// Page
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to convert page to int",
		})
		return

	}

	invoice, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	var shipmentResponses []responses.ShipmentResponse
	for _, invoice := range invoice {
		shipmentResponses = append(shipmentResponses, convertShipmentResponse(invoice))
	}

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: shipmentResponses,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

func (h *ShipmentController) CreateShipment(c *gin.Context) {
	var shipmentForm requests.ShipmentRequest

	err := c.ShouldBindJSON(&shipmentForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shipment, err := h.service.Create(shipmentForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertShipmentResponse(shipment),
	}

	c.JSON(http.StatusOK, webResponse)
}

package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/invoice"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	service services.Service
}

func NewInvoiceController(service services.Service) *InvoiceController {
	return &InvoiceController{service}
}

func convertInvoiceResponse(o models.Invoice) responses.InvoiceResponse {
	var invoiceItemsResponse []responses.InvoiceItemResponse

	for _, item := range o.InvoiceItems {
		itemResponse := responses.InvoiceItemResponse{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     &item.Kurs,
			Quantity: item.Quantity,
		}
		invoiceItemsResponse = append(invoiceItemsResponse, itemResponse)
	}

	return responses.InvoiceResponse{
		ID:            o.ID,
		InvoiceNumber: o.InvoiceNumber,
		Type:          o.Type,
		CustomerID:    o.CustomerID,
		ConsigneeID:   o.ConsigneeID,
		ShipperID:     o.ShipperID,
		Service:       o.Service,
		BLAWB:         o.BLAWB,
		AJU:           o.AJU,
		POL:           o.POL,
		POD:           o.POD,
		ShippingMarks: o.ShippingMarks,
		InvoiceDate:   o.InvoiceDate.Format("2006-01-02"),
		Status:        o.Status,
		InvoiceItems:  invoiceItemsResponse,
	}
}

func (h *InvoiceController) FindAllInvoicesWithoutPagination(c *gin.Context) {
	invoices, err := h.service.FindAllNoPagination()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var invoiceResponses []responses.InvoiceResponse

	if len(invoices) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.InvoiceResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, invoice := range invoices {
		response := convertInvoiceResponse(invoice)

		invoiceResponses = append(invoiceResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   invoiceResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) CreateInvoice(c *gin.Context) {
	var invoiceForm requests.CreateInvoiceRequest

	err := c.ShouldBindJSON(&invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	invoice, err := h.service.Create(invoiceForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindInvoiceByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	invoice, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if invoice.ID == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *InvoiceController) EditInvoice(c *gin.Context) {
	var invoiceForm requests.EditInvoiceRequest

	err := c.ShouldBindJSON(&invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	invoice, err := h.service.Edit(id, invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) DeleteInvoice(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	invoice, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindAll(c *gin.Context) {
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

	var invoiceResponses []responses.InvoiceResponse
	for _, invoice := range invoice {
		invoiceResponses = append(invoiceResponses, convertInvoiceResponse(invoice))
	}

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: invoiceResponses,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

func (h *InvoiceController) GeneratePDF(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	invoice, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if invoice.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Quotation with this ID is not found!",
		})
		return
	}
	filePath := fmt.Sprintf("pdf/invoice/invoice-%s.pdf", sanitizeFilename(invoice.InvoiceNumber))
	// helper.GenerateQuotationPDF(invoice)

	c.File(filePath)
}

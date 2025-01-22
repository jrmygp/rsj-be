package controllers

import (
	"fmt"
	"net/http"
	"server/helper"
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

func convertInvoiceExportResponse(o models.InvoiceExport) responses.InvoiceResponse {
	var invoiceItemsResponse []responses.InvoiceItemResponse
	var nominal float64

	for _, item := range o.InvoiceItems {
		// Convert Quantity to float64 for the calculation
		quantity := float64(item.Quantity)

		// Calculate the subtotal for the current item
		subTotal := item.Price * quantity
		if item.Kurs != nil {
			subTotal *= *item.Kurs // Multiply by Kurs if it's not nil
		}
		nominal += subTotal // Add to the total nominal value

		itemResponse := responses.InvoiceItemResponse{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     item.Kurs,
			Quantity: item.Quantity,
		}
		invoiceItemsResponse = append(invoiceItemsResponse, itemResponse)
	}

	return responses.InvoiceResponse{
		ID:                  o.ID,
		InvoiceNumber:       o.InvoiceNumber,
		Type:                o.Type,
		CustomerID:          o.CustomerID,
		CustomerName:        o.Customer.Name,
		ConsigneeID:         o.ConsigneeID,
		CosgineeName:        o.Consignee.Name,
		ShipperID:           o.ShipperID,
		ShipperName:         o.Shipper.Name,
		Service:             o.Service,
		BLAWB:               o.BLAWB,
		AJU:                 o.AJU,
		PortOfLoadingID:     o.PortOfLoadingID,
		PortOfLoadingName:   o.PortOfLoading.PortName,
		PortOfDischargeID:   o.PortOfDischargeID,
		PortOfDischargeName: o.PortOfDischarge.PortName,
		ShippingMarks:       o.ShippingMarks,
		InvoiceDate:         o.InvoiceDate.Format("2006-01-02"),
		Status:              o.Status,
		InvoiceItems:        invoiceItemsResponse,
		Nominal:             nominal,
	}
}

func convertInvoiceImportResponse(o models.InvoiceImport) responses.InvoiceResponse {
	var invoiceItemsResponse []responses.InvoiceItemResponse
	var nominal float64

	for _, item := range o.InvoiceItems {
		// Convert Quantity to float64 for the calculation
		quantity := float64(item.Quantity)

		// Calculate the subtotal for the current item
		subTotal := item.Price * quantity
		if item.Kurs != nil {
			subTotal *= *item.Kurs // Multiply by Kurs if it's not nil
		}
		nominal += subTotal // Add to the total nominal value

		itemResponse := responses.InvoiceItemResponse{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     item.Kurs,
			Quantity: item.Quantity,
		}
		invoiceItemsResponse = append(invoiceItemsResponse, itemResponse)
	}

	return responses.InvoiceResponse{
		ID:                  o.ID,
		InvoiceNumber:       o.InvoiceNumber,
		Type:                o.Type,
		CustomerID:          o.CustomerID,
		CustomerName:        o.Customer.Name,
		ConsigneeID:         o.ConsigneeID,
		CosgineeName:        o.Consignee.Name,
		ShipperID:           o.ShipperID,
		ShipperName:         o.Shipper.Name,
		Service:             o.Service,
		BLAWB:               o.BLAWB,
		AJU:                 o.AJU,
		PortOfLoadingID:     o.PortOfLoadingID,
		PortOfLoadingName:   o.PortOfLoading.PortName,
		PortOfDischargeID:   o.PortOfDischargeID,
		PortOfDischargeName: o.PortOfDischarge.PortName,
		ShippingMarks:       o.ShippingMarks,
		InvoiceDate:         o.InvoiceDate.Format("2006-01-02"),
		Status:              o.Status,
		InvoiceItems:        invoiceItemsResponse,
		Nominal:             nominal,
	}
}

func convertDoorToDoorResponse(o models.DoorToDoorInvoice) responses.DoorToDoorResponse {
	var invoiceItemsResponse []responses.InvoiceD2DItemResponse
	var nominal float64

	for _, item := range o.InvoiceItems {
		quantity := float64(item.Quantity)

		subTotal := item.Price * quantity
		if item.Kurs != nil {
			subTotal *= *item.Kurs
		}
		nominal += subTotal

		itemResponse := responses.InvoiceD2DItemResponse{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     item.Kurs,
			Quantity: item.Quantity,
		}
		invoiceItemsResponse = append(invoiceItemsResponse, itemResponse)
	}

	return responses.DoorToDoorResponse{
		ID:                  o.ID,
		InvoiceNumber:       o.InvoiceNumber,
		Type:                o.Type,
		CustomerID:          o.CustomerID,
		CustomerName:        o.Customer.Name,
		ConsigneeID:         o.ConsigneeID,
		CosgineeName:        o.Consignee.Name,
		ShipperID:           o.ShipperID,
		ShipperName:         o.Shipper.Name,
		Service:             o.Service,
		PortOfLoadingID:     o.PortOfLoadingID,
		PortOfLoadingName:   o.PortOfLoading.PortName,
		PortOfDischargeID:   o.PortOfDischargeID,
		PortOfDischargeName: o.PortOfDischarge.PortName,
		ShippingMarks:       o.ShippingMarks,
		InvoiceDate:         o.InvoiceDate.Format("2006-01-02"),
		Status:              o.Status,
		Quantity:            o.Quantity,
		Weight:              o.Weight,
		Volume:              o.Volume,
		InvoiceItems:        invoiceItemsResponse,
		Nominal:             nominal,
	}
}

func (h *InvoiceController) FindAllInvoiceExportWithoutPagination(c *gin.Context) {
	invoices, err := h.service.FindAllExportNoPagination()
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
		response := convertInvoiceExportResponse(invoice)

		invoiceResponses = append(invoiceResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   invoiceResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) CreateInvoiceExport(c *gin.Context) {
	var invoiceForm requests.CreateInvoiceRequest

	err := c.ShouldBindJSON(&invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	invoice, err := h.service.CreateExport(invoiceForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceExportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindInvoiceExportByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	invoice, err := h.service.FindExportByID(ID)
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
		Data:   convertInvoiceExportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *InvoiceController) EditInvoiceExport(c *gin.Context) {
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

	// Get user role from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "User information missing",
		})
		return
	}
	userRoleID := int(user.(models.User).UserRoleID)

	invoice, err := h.service.EditExport(id, invoiceForm, userRoleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceExportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) DeleteInvoiceExport(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	invoice, err := h.service.DeleteExport(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceExportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindAllExport(c *gin.Context) {
	var filterForm requests.InvoiceFilterRequest

	err := c.ShouldBindJSON(&filterForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	invoice, totalCount, firstRow, lastRow, totalPages := h.service.FindAllExport(searchQuery, page, filterForm)

	var invoiceResponses []responses.InvoiceResponse
	for _, invoice := range invoice {
		invoiceResponses = append(invoiceResponses, convertInvoiceExportResponse(invoice))
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

	invoice, err := h.service.FindExportByID(ID)
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
	filePath := fmt.Sprintf("pdf/invoice/%s.pdf", sanitizeFilename(invoice.InvoiceNumber))
	helper.GenerateInvoicePDF(invoice)

	c.File(filePath)
}

func (h *InvoiceController) FindAllInvoiceImportWithoutPagination(c *gin.Context) {
	invoices, err := h.service.FindAllImportNoPagination()
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
		response := convertInvoiceImportResponse(invoice)

		invoiceResponses = append(invoiceResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   invoiceResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) CreateInvoiceImport(c *gin.Context) {
	var invoiceForm requests.CreateInvoiceRequest

	err := c.ShouldBindJSON(&invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	invoice, err := h.service.CreateImport(invoiceForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceImportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindInvoiceImportByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	invoice, err := h.service.FindImportByID(ID)
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
		Data:   convertInvoiceImportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *InvoiceController) EditInvoiceImport(c *gin.Context) {
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

	// Get user role from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "User information missing",
		})
		return
	}
	userRoleID := int(user.(models.User).UserRoleID)

	invoice, err := h.service.EditImport(id, invoiceForm, userRoleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceImportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) DeleteInvoiceImport(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	invoice, err := h.service.DeleteImport(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertInvoiceImportResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindAllImport(c *gin.Context) {
	var filterForm requests.InvoiceFilterRequest

	err := c.ShouldBindJSON(&filterForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	invoice, totalCount, firstRow, lastRow, totalPages := h.service.FindAllImport(searchQuery, page, filterForm)

	var invoiceResponses []responses.InvoiceResponse
	for _, invoice := range invoice {
		invoiceResponses = append(invoiceResponses, convertInvoiceImportResponse(invoice))
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

// Door to Door controllers
func (h *InvoiceController) FindAllDoorToDoorWithoutPagination(c *gin.Context) {
	invoices, err := h.service.FindAllDoorToDoorNoPagination()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var invoiceResponses []responses.DoorToDoorResponse

	if len(invoices) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.DoorToDoorResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, invoice := range invoices {
		response := convertDoorToDoorResponse(invoice)

		invoiceResponses = append(invoiceResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   invoiceResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) CreateDoorToDoor(c *gin.Context) {
	var invoiceForm requests.CreateDoorToDoorRequest

	err := c.ShouldBindJSON(&invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	invoice, err := h.service.CreateDoorToDoor(invoiceForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertDoorToDoorResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindDoorToDoorByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	invoice, err := h.service.FindDoorToDoorByID(ID)
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
		Data:   convertDoorToDoorResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *InvoiceController) EditDoorToDoor(c *gin.Context) {
	var invoiceForm requests.EditDoorToDoorRequest

	err := c.ShouldBindJSON(&invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	invoice, err := h.service.EditDoorToDoor(id, invoiceForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertDoorToDoorResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) DeleteDoorToDoor(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	invoice, err := h.service.DeleteDoorToDoor(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertDoorToDoorResponse(invoice),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *InvoiceController) FindAllDoorToDoor(c *gin.Context) {
	var filterForm requests.InvoiceFilterRequest

	err := c.ShouldBindJSON(&filterForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	invoice, totalCount, firstRow, lastRow, totalPages := h.service.FindAllDoorToDoor(searchQuery, page, filterForm)

	var invoiceResponses []responses.DoorToDoorResponse
	for _, invoice := range invoice {
		invoiceResponses = append(invoiceResponses, convertDoorToDoorResponse(invoice))
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

func (h *InvoiceController) GenerateDoorToDoorPDF(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	invoice, err := h.service.FindDoorToDoorByID(ID)
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
	filePath := fmt.Sprintf("pdf/invoice-d2d/%s.pdf", sanitizeFilename(invoice.InvoiceNumber))
	helper.GenerateInvoiceD2DPDF(invoice)

	c.File(filePath)
}

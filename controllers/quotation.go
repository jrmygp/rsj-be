package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/quotation"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuotationController struct {
	service services.Service
}

func NewQuotationController(service services.Service) *QuotationController {
	return &QuotationController{service}
}

func convertQuotationResponse(o models.Quotation) responses.QuotationResponse {
	var listChargesResponse []responses.ChargeResponse

	for _, charge := range o.ListCharges {
		chargeResponse := responses.ChargeResponse{
			ItemName: charge.ItemName,
			Price:    charge.Price,
			Currency: charge.Currency,
			RatioIDR: charge.RatioToIDR,
			Quantity: charge.Quantity,
			Unit:     charge.Unit,
			Note:     charge.Note,
		}
		listChargesResponse = append(listChargesResponse, chargeResponse)
	}

	return responses.QuotationResponse{
		QuotationNumber:   o.QuotationNumber,
		RateValidity:      o.RateValidity,
		ShippingTerm:      o.ShippingTerm,
		Service:           o.Service,
		Status:            o.Status,
		Commodity:         o.Commodity,
		Weight:            o.Weight,
		Volume:            o.Volume,
		Note:              o.Note,
		SalesID:           o.SalesID,
		CustomerID:        o.CustomerID,
		PortOfLoadingID:   o.PortOfLoadingID,
		PortOfDischargeID: o.PortOfDischargeID,
		ListCharges:       listChargesResponse,
	}
}

func (h *QuotationController) CreateQuotation(c *gin.Context) {
	var quotationForm requests.CreateQuotationRequest

	err := c.ShouldBindJSON(&quotationForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	quotation, err := h.service.Create(quotationForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertQuotationResponse(quotation),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *QuotationController) FindQuotationByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	quotation, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if quotation.ID == 0 {
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
		Data:   convertQuotationResponse(quotation),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *QuotationController) EditQuotation(c *gin.Context) {
	var quotationForm requests.EditQuotationRequest

	err := c.ShouldBindJSON(&quotationForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	quotation, err := h.service.Edit(id, quotationForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertQuotationResponse(quotation),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *QuotationController) DeleteQuotation(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	quotation, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertQuotationResponse(quotation),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *QuotationController) FindAll(c *gin.Context) {
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

	quotation, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: quotation,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

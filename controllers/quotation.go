package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/quotation"

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
			Note:     &charge.Note,
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

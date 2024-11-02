package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/costCharges"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CostChargesController struct {
	service services.Service
}

func NewCostChargesController(service services.Service) *CostChargesController {
	return &CostChargesController{service}
}

func convertCostChargesResponse(o models.CostCharges) responses.CostChargesResponse {
	return responses.CostChargesResponse{
		ID:     o.ID,
		Name:   o.Name,
		Status: o.Status,
	}
}

func (h *CostChargesController) CreateCostCharge(c *gin.Context) {
	var costChargeForm requests.CreateCostChargesRequest

	err := c.ShouldBindJSON(&costChargeForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	costCharge, err := h.service.Create(costChargeForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCostChargesResponse(costCharge),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CostChargesController) FindCostChargeByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	costCharge, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if costCharge.ID == 0 {
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
		Data:   convertCostChargesResponse(costCharge),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CostChargesController) EditCostCharge(c *gin.Context) {
	var costChargeForm requests.EditCostChargesRequest

	err := c.ShouldBindJSON(&costChargeForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	costCharge, err := h.service.Edit(id, costChargeForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCostChargesResponse(costCharge),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CostChargesController) DeleteCostCharge(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	costCharge, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCostChargesResponse(costCharge),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CostChargesController) FindAll(c *gin.Context) {
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

	costCharge, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: costCharge,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

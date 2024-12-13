package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/shipper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShipperController struct {
	service services.Service
}

func NewShipperController(service services.Service) *ShipperController {
	return &ShipperController{service}
}

func convertShipperResponse(o models.Shipper) responses.ShipperResponse {
	return responses.ShipperResponse{
		ID:      o.ID,
		Name:    o.Name,
		Address: o.Address,
	}
}

func (h *ShipperController) FindAllShippersWithoutPagination(c *gin.Context) {
	shippers, err := h.service.FindAllNoPagination()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var shipperResponses []responses.ShipperResponse

	if len(shippers) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.ShipperResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, shipper := range shippers {
		response := convertShipperResponse(shipper)

		shipperResponses = append(shipperResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   shipperResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *ShipperController) CreateShipper(c *gin.Context) {
	var shipperForm requests.CreateShipperRequest

	err := c.ShouldBindJSON(&shipperForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shipper, err := h.service.Create(shipperForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertShipperResponse(shipper),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *ShipperController) FindShipperByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	shipper, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if shipper.ID == 0 {
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
		Data:   convertShipperResponse(shipper),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *ShipperController) EditShipper(c *gin.Context) {
	var shipperForm requests.EditShipperRequest

	err := c.ShouldBindJSON(&shipperForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	shipper, err := h.service.Edit(id, shipperForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertShipperResponse(shipper),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *ShipperController) DeleteShipper(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	shipper, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertShipperResponse(shipper),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *ShipperController) FindAll(c *gin.Context) {
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

	shippers, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: shippers,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

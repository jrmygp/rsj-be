package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/warehouse"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WarehouseController struct {
	service services.Service
}

func NewWarehouseController(service services.Service) *WarehouseController {
	return &WarehouseController{service}
}

func convertWarehouseResponse(o models.Warehouse) responses.WarehouseResponse {
	return responses.WarehouseResponse{
		ID:         o.ID,
		Name:       o.Name,
		Code:       o.Code,
		FlightName: o.FlightName,
		FlightCode: o.FlightCode,
	}
}

func (h *WarehouseController) FindAllWarehousesWithoutPagination(c *gin.Context) {
	warehouses, err := h.service.FindAllNoPagination()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var warehouseResponses []responses.WarehouseResponse

	if len(warehouses) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.WarehouseResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, warehouse := range warehouses {
		response := convertWarehouseResponse(warehouse)

		warehouseResponses = append(warehouseResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   warehouseResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *WarehouseController) CreateWarehouse(c *gin.Context) {
	var warehouseForm requests.WarehouseRequest

	err := c.ShouldBindJSON(&warehouseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	warehouse, err := h.service.Create(warehouseForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertWarehouseResponse(warehouse),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *WarehouseController) FindWarehouseByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	warehouse, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if warehouse.ID == 0 {
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
		Data:   convertWarehouseResponse(warehouse),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *WarehouseController) EditWarehouse(c *gin.Context) {
	var warehouseForm requests.WarehouseRequest

	err := c.ShouldBindJSON(&warehouseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	warehouse, err := h.service.Edit(id, warehouseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertWarehouseResponse(warehouse),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *WarehouseController) DeleteWarehouse(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	warehouse, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertWarehouseResponse(warehouse),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *WarehouseController) FindAll(c *gin.Context) {
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

	warehouse, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: warehouse,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

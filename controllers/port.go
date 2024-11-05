package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/port"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PortController struct {
	service services.Service
}

func NewPortController(service services.Service) *PortController {
	return &PortController{service}
}

func convertPortResponse(o models.Port) responses.PortResponse {
	return responses.PortResponse{
		ID:       o.ID,
		PortName: o.PortName,
		Note:     o.Note,
		Status:   o.Status,
	}
}

func (h *PortController) FindAllPortsWithoutPagination(c *gin.Context) {
	ports, err := h.service.FindAllNoPagination()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var portResponses []responses.PortResponse

	if len(ports) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.PortResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, port := range ports {
		response := convertPortResponse(port)

		portResponses = append(portResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   portResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *PortController) CreatePort(c *gin.Context) {
	var portForm requests.CreatePortRequest

	err := c.ShouldBindJSON(&portForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	port, err := h.service.Create(portForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertPortResponse(port),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *PortController) FindPortByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	port, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if port.ID == 0 {
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
		Data:   convertPortResponse(port),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *PortController) EditPort(c *gin.Context) {
	var portForm requests.EditPortRequest

	err := c.ShouldBindJSON(&portForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	port, err := h.service.Edit(id, portForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertPortResponse(port),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *PortController) DeletePort(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	port, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertPortResponse(port),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *PortController) FindAll(c *gin.Context) {
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

	port, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: port,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

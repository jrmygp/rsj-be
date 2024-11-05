package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/customer"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service services.Service
}

func NewCustomerController(service services.Service) *CustomerController {
	return &CustomerController{service}
}

func convertCustomerResponse(o models.Customer) responses.CustomerResponse {
	return responses.CustomerResponse{
		ID:      o.ID,
		Name:    o.Name,
		Address: o.Address,
	}
}

func (h *CustomerController) FindAllCustomersWithoutPagination(c *gin.Context) {
	customers, err := h.service.FindAllNoPagination()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var customerResponses []responses.CustomerResponse

	if len(customers) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.CustomerResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, customer := range customers {
		response := convertCustomerResponse(customer)

		customerResponses = append(customerResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   customerResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CustomerController) CreateCustomer(c *gin.Context) {
	var customerForm requests.CreateCustomerRequest

	err := c.ShouldBindJSON(&customerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer, err := h.service.Create(customerForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCustomerResponse(customer),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CustomerController) FindCustomerByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	customer, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if customer.ID == 0 {
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
		Data:   convertCustomerResponse(customer),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CustomerController) EditCustomer(c *gin.Context) {
	var customerForm requests.EditCustomerRequest

	err := c.ShouldBindJSON(&customerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	customer, err := h.service.Edit(id, customerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCustomerResponse(customer),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CustomerController) DeleteCustomer(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	customer, err := h.service.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCustomerResponse(customer),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *CustomerController) FindAll(c *gin.Context) {
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

	customers, totalCount, firstRow, lastRow, totalPages := h.service.FindAll(searchQuery, page)

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: customers,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

package controllers

import (
	"fmt"
	"net/http"
	"server/helper"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/document"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	service services.Service
}

func NewDocumentController(service services.Service) *DocumentController {
	return &DocumentController{service}
}

func convertSuratTugasResponse(o models.SuratTugas) responses.SuratTugasResponse {
	return responses.SuratTugasResponse{
		ID:             o.ID,
		DocumentNumber: o.DocumentNumber,
		Assignor:       o.Assignor,
		Assignee:       o.Assignee,
		Liners:         o.Liners,
		Type:           o.Type,
		BLAWB:          o.BLAWB,
		Date:           o.Date.Format("2006-01-02"),
	}
}

func convertSuratJalanResponse(o models.SuratJalan) responses.SuratJalanResponse {
	var itemsResponse []responses.ItemResponse

	for _, item := range o.Items {
		itemResponse := responses.ItemResponse{
			ItemName: item.ItemName,
			Quantity: item.Quantity,
			Colly:    item.Colly,
			Volume:   item.Volume,
			Unit:     item.Unit,
			Note:     item.Note,
		}
		itemsResponse = append(itemsResponse, itemResponse)
	}

	return responses.SuratJalanResponse{
		ID:             o.ID,
		DocumentNumber: o.DocumentNumber,
		Recipient:      o.Recipient,
		Address:        o.Address,
		Date:           o.DocumentNumber,
		Items:          itemsResponse,
	}
}

// Surat Tugas
func (h *DocumentController) FindAllSuratTugas(c *gin.Context) {
	searchQuery := c.Query("search")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to convert page to int",
		})
		return
	}

	document, totalCount, firstRow, lastRow, totalPages := h.service.FindAllSuratTugas(searchQuery, page)

	var documentResponses []responses.SuratTugasResponse
	for _, document := range document {
		documentResponses = append(documentResponses, convertSuratTugasResponse(document))
	}

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: documentResponses,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

func (h *DocumentController) CreateSuratTugas(c *gin.Context) {
	var suratTugas requests.CreateSuratTugasRequest

	err := c.ShouldBindJSON(&suratTugas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	document, err := h.service.CreateSuratTugas(suratTugas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSuratTugasResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *DocumentController) FindSuratTugasByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	document, err := h.service.FindSuratTugasByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if document.ID == 0 {
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
		Data:   convertSuratTugasResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *DocumentController) EditSuratTugas(c *gin.Context) {
	var suratTugas requests.EditSuratTugasRequest

	err := c.ShouldBindJSON(&suratTugas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	document, err := h.service.EditSuratTugas(id, suratTugas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSuratTugasResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *DocumentController) DeleteSuratTugas(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	document, err := h.service.DeleteSuratTugas(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSuratTugasResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *DocumentController) GenerateSuratTugasPDF(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	document, err := h.service.FindSuratTugasByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if document.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Surat Tugas with this ID is not found!",
		})
		return
	}
	filePath := fmt.Sprintf("pdf/surat-tugas/%s.pdf", sanitizeFilename(document.DocumentNumber))
	helper.GenerateSuratTugasPDF(document)

	c.File(filePath)
}

// Surat Jalan
func (h *DocumentController) FindAllSuratJalan(c *gin.Context) {
	searchQuery := c.Query("search")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to convert page to int",
		})
		return
	}

	document, totalCount, firstRow, lastRow, totalPages := h.service.FindAllSuratJalan(searchQuery, page)

	var documentResponses []responses.SuratJalanResponse
	for _, document := range document {
		documentResponses = append(documentResponses, convertSuratJalanResponse(document))
	}

	webPaginationResponse := responses.PaginationResponse{
		Code:          http.StatusOK,
		Status:        "OK",
		DataResponses: documentResponses,
		TotalCount:    totalCount,
		FirstRow:      firstRow,
		LastRow:       lastRow,
		TotalPages:    totalPages,
	}

	c.JSON(http.StatusOK, webPaginationResponse)
}

func (h *DocumentController) CreateSuratJalan(c *gin.Context) {
	var suratJalan requests.CreateSuratJalanRequest

	err := c.ShouldBindJSON(&suratJalan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	document, err := h.service.CreateSuratJalan(suratJalan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSuratJalanResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *DocumentController) FindSuratJalanByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	document, err := h.service.FindSuratJalanByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if document.ID == 0 {
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
		Data:   convertSuratJalanResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *DocumentController) EditSuratJalan(c *gin.Context) {
	var suratJalan requests.EditSuratJalanRequest

	err := c.ShouldBindJSON(&suratJalan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	document, err := h.service.EditSuratJalan(id, suratJalan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSuratJalanResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *DocumentController) DeleteSuratJalan(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	document, err := h.service.DeleteSuratJalan(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSuratJalanResponse(document),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *DocumentController) GenerateSuratJalanPDF(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	document, err := h.service.FindSuratJalanByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if document.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Surat Jalan with this ID is not found!",
		})
		return
	}
	filePath := fmt.Sprintf("pdf/surat-jalan/%s.pdf", sanitizeFilename(document.DocumentNumber))
	helper.GenerateSuratJalanPDF(document)

	c.File(filePath)
}

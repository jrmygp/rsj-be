package requests

import "server/helper"

type CreateSuratTugasRequest struct {
	DocumentNumber string            `json:"documentNumber" binding:"required"`
	Assignor       string            `json:"assignor" binding:"required"`
	Assignee       string            `json:"assignee" binding:"required"`
	Liners         string            `json:"liners" binding:"required"`
	Type           string            `json:"type" binding:"required"`
	BLAWB          string            `json:"blawb" binding:"required"`
	Date           helper.CustomDate `json:"date" binding:"required"`
}

type EditSuratTugasRequest struct {
	DocumentNumber string            `json:"documentNumber"`
	Assignor       string            `json:"assignor"`
	Assignee       string            `json:"assignee"`
	Liners         string            `json:"liners"`
	Type           string            `json:"type"`
	BLAWB          string            `json:"blawb"`
	Date           helper.CustomDate `json:"date"`
}

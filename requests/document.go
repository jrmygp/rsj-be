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

type ItemRequest struct {
	ItemName string   `json:"itemName" binding:"required"`
	Quantity int      `json:"quantity" binding:"required"`
	Colly    *float64 `json:"colly"`
	Volume   *float64 `json:"volume"`
	Unit     string   `json:"unit" binding:"required"`
	Note     *string  `json:"note"`
}

type CreateSuratJalanRequest struct {
	DocumentNumber string            `json:"documentNumber" binding:"required"`
	Recipient      string            `json:"recipient" binding:"required"`
	Address        string            `json:"address" binding:"required"`
	Date           helper.CustomDate `json:"date" binding:"required"`
	Items          []ItemRequest     `json:"items" binding:"required"`
}

type EditSuratJalanRequest struct {
	DocumentNumber string            `json:"documentNumber"`
	Recipient      string            `json:"recipient"`
	Address        string            `json:"address"`
	Date           helper.CustomDate `json:"date"`
	Items          []ItemRequest     `json:"items"`
}

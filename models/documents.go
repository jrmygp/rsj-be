package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type SuratTugas struct {
	gorm.Model
	ID             int
	DocumentNumber string
	Assignor       string
	Assignee       string
	Liners         string
	Type           string
	BLAWB          string
	Date           time.Time
}

type Item struct {
	ItemName string   `json:"itemName"`
	Quantity int      `json:"quantity"`
	Colly    *float64 `json:"colly"`
	Volume   *float64 `json:"volume"`
	Unit     string   `json:"unit"`
	Note     *string  `json:"note"`
}

type JSONItems []Item

func (js JSONItems) Value() (driver.Value, error) {
	return json.Marshal(js)
}

func (js *JSONItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, js)
}

type SuratJalan struct {
	gorm.Model
	ID             int
	DocumentNumber string
	Recipient      string
	Address        string
	Date           time.Time
	Items          JSONItems
}

package responses

type SuratTugasResponse struct {
	ID             int    `json:"id"`
	DocumentNumber string `json:"documentNumber"`
	Assignor       string `json:"assignor"`
	Assignee       string `json:"assignee"`
	Liners         string `json:"liners"`
	Type           string `json:"type"`
	BLAWB          string `json:"blawb"`
	Date           string `json:"date"`
}

type ItemResponse struct {
	ItemName string   `json:"itemName"`
	Quantity int      `json:"quantity"`
	Colly    *float64 `json:"colly"`
	Volume   *float64 `json:"volume"`
	Unit     string   `json:"unit"`
	Note     *string  `json:"note"`
}

type SuratJalanResponse struct {
	ID             int            `json:"id"`
	DocumentNumber string         `json:"documentNumber"`
	Recipient      string         `json:"recipient"`
	Address        string         `json:"address"`
	Date           string         `json:"date"`
	Items          []ItemResponse `json:"items"`
}

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

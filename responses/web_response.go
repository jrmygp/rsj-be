package responses

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type PaginationResponse struct {
	Code          int         `json:"code"`
	Status        string      `json:"status"`
	DataResponses interface{} `json:"data"`
	TotalCount    int64       `json:"totalCount"`
	FirstRow      int         `json:"firstRow"`
	LastRow       int         `json:"lastRow"`
	TotalPages    int         `json:"totalPages"`
}

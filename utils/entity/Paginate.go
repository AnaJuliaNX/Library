package entity

type Meta struct {
	Total       int64 `json:"total"`
	CurrentPage int64 `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

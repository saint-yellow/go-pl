package serialization

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Page int `json:"page"`
	Total int `json:"total"`
}

type PaginationResult struct {
	Result

	Page int `json:"page"`
	Total int `json:"total"`
}

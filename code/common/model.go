package common

// CommonRequest 通用请求体
type CommonRequest[T any] struct {
	Param T `json:"param"`
}

// CommonResponse 通用响应体
type CommonResponse[T any] struct {
	Data    T      `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// PageRequest 通用分页请求体
type PageRequest[T any] struct {
	Param     T   `json:"param"`
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

// PageResponse 通用分页响应体
type PageResponse[T any] struct {
	Data      []T    `json:"data"`
	PageSize  int    `json:"pageSize"`
	PageIndex int    `json:"pageIndex"`
	TotalRow  int    `json:"totalRow"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

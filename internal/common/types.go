package common

type Currency string
type Id uint

type GetAllDTO[T any] struct {
	Data          T   `json:"data"`
	PageNumber    int `json:"pageNumber"`
	PageSize      int `json:"pageSize"`
	TotalRowCount int `json:"totalRowCount"`
}

type GetQueryParams struct {
	PageSize   *int
	PageNumber *int
	Fields     *[]string
	Embed      *string
}

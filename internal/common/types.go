package common

type EID uint

type BaseEntity struct {
	Id EID
}

type GetAllDTO[T any] struct {
	Data          T   `json:"data"`
	PageNumber    int `json:"pageNumber"`
	PageSize      int `json:"pageSize"`
	TotalRowCount int `json:"totalRowCount"`
}

type GetQueryParams struct {
	PageSize   *int      `query:"page-size"`
	PageNumber *int      `query:"page"`
	Fields     *[]string `query:"fields"`
	Embed      *string   `query:"embed"`
}

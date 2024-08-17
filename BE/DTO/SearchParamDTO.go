package DTO

type ListDTOModel[T any] struct {
	PageIndex   int64
	PageSize    int64
	TotalCount  int64
	TotalNotify int64
	CalValue1   float64
	CalValue2   float64
	CalValue3   float64
	Source      *[]T
}

type SearchParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddedParam map[string]string

type GetSearchParam struct {
	PageIndex    int64          `json:"pageIndex"`
	PageSize     int64          `json:"pageSize"`
	SearchParams *[]SearchParam `json:"searchParams"`
}

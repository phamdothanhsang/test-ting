package ViewModels

import "wan-api-kol-event/DTO"

type KolViewModel struct {
	Result       string        `json:"result"`       //Result : success, unsuccess
	ErrorMessage string        `json:"errorMessage"` // The query error
	PageIndex    int64         `json:"pageIndex"`    //The page index
	PageSize     int64         `json:"pageSize"`     //The page size
	Guid         string        `json:"guid"`
	TotalCount   int64         `json:"totalCount"`
	KOL          []*DTO.KolDTO `json:"kol"`
}

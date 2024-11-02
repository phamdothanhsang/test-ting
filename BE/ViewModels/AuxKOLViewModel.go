package ViewModels

import (
	"encoding/json"
	"wan-api-kol-event/DTO"
)

type AuxKolViewModel struct {
	Result       string        `json:"result"`       //Result : success, unsuccess
	ErrorMessage string        `json:"errorMessage"` // The query error
	PageIndex    int64         `json:"pageIndex"`    //The page index
	PageSize     int64         `json:"pageSize"`     //The page size
	Guid         string        `json:"-"`
	TotalCount   int64         `json:"totalCount"`
	KOL          []*DTO.KolDTO `json:"KolInformation"`
}


func (m KolViewModel) MarshalJSON() ([]byte, error) {
	auxKolViewModel := (*AuxKolViewModel)(&m)
	return json.Marshal(auxKolViewModel)
}
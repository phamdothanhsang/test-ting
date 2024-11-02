package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex int, pageSize int) ([]*DTO.KolDTO, error) {
	var kols []*Models.Kol
	var kolDTOs []*DTO.KolDTO
	result := Initializers.DB.
				Offset(((pageIndex - 1) * pageSize)).
				Limit(pageSize).
				Find(&kols)
				
	for _, kol := range kols {
		kolDTO := (*DTO.KolDTO)(kol)
		kolDTOs = append(kolDTOs, kolDTO)
	}
	return kolDTOs, result.Error
}

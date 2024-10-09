package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
)

var KOLs []*DTO.KolDTO
var DB = Initializers.DB

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex int, pageSize int) ([]*DTO.KolDTO, error) {
	result := DB.Find(&KOLs)

	return KOLs, result.Error
}

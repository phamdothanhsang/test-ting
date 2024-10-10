package Logic

import (
	"fmt"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
)

var KOLs []*DTO.KolDTO

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageSize int, pageIndex int) ([]*DTO.KolDTO, error) {
	var DB = Initializers.DB

	if pageIndex <= 0 {
		pageIndex = 1
	}

	switch {
	case pageSize > 20:
		pageSize = 20
	case pageSize < 0:
		pageSize = 5
	}

	offset := (pageIndex - 1) * pageSize
	fmt.Println(offset, pageSize)

	result := DB.Offset(offset).Limit(pageSize).Find(&KOLs)

	return KOLs, result.Error
}

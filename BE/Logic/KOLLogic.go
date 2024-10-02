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
func GetKolLogic(pageIndex int64, pageSize int64) ([]*DTO.KolDTO, error) {
	var kols []Models.Kol

	offset := (pageIndex - 1) * pageSize

	if err := Initializers.DB.Limit(int(pageSize)).Offset(int(offset)).Find(&kols).Error; err != nil {
		return nil, err
	}

	var kolDTOs []*DTO.KolDTO

	for _, kol := range kols {
		kolDTOs = append(kolDTOs, DTO.NewKolDTO(kol))
	}
	return kolDTOs, nil
}

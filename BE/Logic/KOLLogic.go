package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

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

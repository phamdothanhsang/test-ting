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
	var kols []Models.Kol

	// USE GORM TO QUERY THE DATABASE
	offset := (pageIndex - 1) * pageSize
	result := Initializers.DB.Debug().Limit(pageSize).Offset(offset).Find(&kols)

	if result.Error != nil {
		return nil, result.Error
	}

	// Model => DTO
	var kolDTOs []*DTO.KolDTO
	for _, kol := range kols {
		kolDTO := &DTO.KolDTO{
			KolID:                kol.KolID,
			UserProfileID:        kol.UserProfileID,
			Language:             kol.Language,
			Education:            kol.Education,
			ExpectedSalary:       kol.ExpectedSalary,
			ExpectedSalaryEnable: kol.ExpectedSalaryEnable,
			ChannelSettingTypeID: kol.ChannelSettingTypeID,
			IDFrontURL:           kol.IDFrontURL,
			IDBackURL:            kol.IDBackURL,
			PortraitURL:          kol.PortraitURL,
			RewardID:             kol.RewardID,
			PaymentMethodID:      kol.PaymentMethodID,
			TestimonialsID:       kol.TestimonialsID,
			VerificationStatus:   kol.VerificationStatus,
			Enabled:              kol.Enabled,
			ActiveDate:           kol.ActiveDate,
			Active:               kol.Active,
			CreatedBy:            kol.CreatedBy,
			CreatedDate:          kol.CreatedDate,
			ModifiedBy:           kol.ModifiedBy,
			ModifiedDate:         kol.ModifiedDate,
			IsRemove:             kol.IsRemove,
			IsOnBoarding:         kol.IsOnBoarding,
			Code:                 kol.Code,
			PortraitRightURL:     kol.PortraitRightURL,
			PortraitLeftURL:      kol.PortraitLeftURL,
			LivenessStatus:       kol.LivenessStatus,
		}
		kolDTOs = append(kolDTOs, kolDTO)
	}

	return kolDTOs, nil
}

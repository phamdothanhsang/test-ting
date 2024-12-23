package Logic

import (
	"errors"

	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func GetKolLogic(pageIndex, pageSize int) ([]*DTO.KolDTO, int64, error) {
	var kols []Models.Kol
	var totalCount int64

	// Count the total number of KOLs
	err := Initializers.DB.Model(&Models.Kol{}).Count(&totalCount).Error
	if err != nil {
		return nil, 0, errors.New("Error counting total KOLs: " + err.Error())
	}

	// Query with pagination using LIMIT and OFFSET
	result := Initializers.DB.Debug().
		Table("\"KOL\"").
		Limit(pageSize).
		Offset(pageIndex * pageSize).
		Find(&kols)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Map the data to DTO
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

	return kolDTOs, totalCount, nil
}

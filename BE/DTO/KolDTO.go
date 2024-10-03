package DTO

import (
	"time"
	"wan-api-kol-event/Models"
)

type KolDTO struct {
	KolID                int64     `json:"kolID"`
	UserProfileID        int64     `json:"userProfileID"`
	Language             string    `json:"language"`
	Education            string    `json:"education"`
	ExpectedSalary       int64     `json:"expectedSalary"`
	ExpectedSalaryEnable bool      `json:"expectedSalaryEnable"`
	ChannelSettingTypeID int64     `json:"channelSettingTypeID"`
	IDFrontURL           string    `json:"iDFrontURL"`
	IDBackURL            string    `json:"iDBackURL"`
	PortraitURL          string    `json:"portraitURL"`
	RewardID             int64     `json:"rewardID"`
	PaymentMethodID      int64     `json:"paymentMethodID"`
	TestimonialsID       int64     `json:"testimonialsID"`
	VerificationStatus   string    `json:"verificationStatus"`
	Enabled              bool      `json:"enabled"`
	ActiveDate           time.Time `json:"activeDate"`
	Active               bool      `json:"active"`
	CreatedBy            string    `json:"createdBy"`
	CreatedDate          time.Time `json:"createdDate"`
	ModifiedBy           string    `json:"modifiedBy"`
	ModifiedDate         time.Time `json:"modifiedDate"`
	IsRemove             bool      `json:"isRemove"`
	IsOnBoarding         bool      `json:"isOnBoarding"`
	Code                 string    `json:"code"`
	PortraitRightURL     string    `json:"portraitRightURL"`
	PortraitLeftURL      string    `json:"portraitLeftURL"`
	LivenessStatus       string    `json:"livenessStatus"`
}

func NewKolDTO(kol Models.Kol) *KolDTO {
	return &KolDTO{
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
		VerificationStatus:   mapVerificationStatus(kol.VerificationStatus),
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
		LivenessStatus:       mapLivenessStatus(kol.LivenessStatus),
	}
}

func mapVerificationStatus(status bool) string {
	if status {
		return "Verified"
	}
	return "Pending"
}

func mapLivenessStatus(status bool) string {
	if status {
		return "Passed"
	}
	return "Failed"
}

package DTO

import (
	"encoding/json"
	"time"
)

type AuxKolDTO struct {
	KolID                int64     `json:"KolID"`
	UserProfileID        int64     `json:"UserProfileID"`
	Language             string    `json:"Language"`
	Education            string    `json:"Education"`
	ExpectedSalary       int64     `json:"ExpectedSalary"`
	ExpectedSalaryEnable bool      `json:"ExpectedSalaryEnable"`
	ChannelSettingTypeID int64     `json:"ChannelSettingTypeID"`
	IDFrontURL           string    `json:"IDFrontURL"`
	IDBackURL            string    `json:"IDBackURL"`
	PortraitURL          string    `json:"PortraitURL"`
	RewardID             int64     `json:"RewardID"`
	PaymentMethodID      int64     `json:"PaymentMethodID"`
	TestimonialsID       int64     `json:"TestimonialsID"`
	VerificationStatus   string    `json:"VerificationStatus"`
	Enabled              bool      `json:"Enabled"`
	ActiveDate           time.Time `json:"ActiveDate"`
	Active               bool      `json:"Active"`
	CreatedBy            string    `json:"CreatedBy"`
	CreatedDate          time.Time `json:"CreatedDate"`
	ModifiedBy           string    `json:"ModifiedBy"`
	ModifiedDate         time.Time `json:"ModifiedDate"`
	IsRemove             bool      `json:"IsRemove"`
	IsOnBoarding         bool      `json:"IsOnBoarding"`
	Code                 string    `json:"Code"`
	PortraitRightURL     string    `json:"PortraitRightURL"`
	PortraitLeftURL      string    `json:"PortraitLeftURL"`
	LivenessStatus       string    `json:"LivenessStatus"`
}

func (m KolDTO) MarshalJSON() ([]byte, error) {
	
	var verificationStatus string
	var livenessStatus string
	if m.VerificationStatus {
		verificationStatus = "Verified"
	} else {
		verificationStatus = "Pending"
	}
	if m.LivenessStatus {
		livenessStatus = "Passed"
	} else {
		livenessStatus = "Failed"
	}
	return json.Marshal(
		&AuxKolDTO{
			KolID: m.KolID,
			UserProfileID: m.UserProfileID,
			Language: m.Language,
			Education: m.Education,
			ExpectedSalary: m.ExpectedSalary,
			ExpectedSalaryEnable: m.ExpectedSalaryEnable,
			ChannelSettingTypeID: m.ChannelSettingTypeID,
			IDFrontURL: m.IDFrontURL,
			IDBackURL: m.IDBackURL,
			PortraitURL: m.PortraitURL,
			RewardID: m.RewardID,
			PaymentMethodID: m.PaymentMethodID,
			TestimonialsID: m.TestimonialsID,
			VerificationStatus: verificationStatus,
			Enabled: m.Enabled,
			ActiveDate: m.ActiveDate,
			Active: m.Active,
			CreatedBy: m.CreatedBy,
			CreatedDate: m.CreatedDate,
			ModifiedBy: m.ModifiedBy,
			ModifiedDate: m.ModifiedDate,
			IsRemove: m.IsRemove,
			IsOnBoarding: m.IsOnBoarding,
			Code: m.Code,
			PortraitRightURL: m.PortraitRightURL,
			PortraitLeftURL: m.PortraitLeftURL,
			LivenessStatus: livenessStatus,
		},
	)
}
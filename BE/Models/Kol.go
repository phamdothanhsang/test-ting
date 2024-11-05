package Models

import (
	"time"
	"wan-api-kol-event/Const"
)

type Kol struct {
	KolID                int64     `json:"kolID"  gorm:"primaryKey;  column:kolid"`
	UserProfileID        int64     `json:"userProfileID" gorm:"column:userprofileid"`
	Language             string    `json:"language" gorm:"column:language"`
	Education            string    `json:"education" gorm:"column:education"`
	ExpectedSalary       int64     `json:"expectedSalary" gorm:"column:expectedsalary"`
	ExpectedSalaryEnable bool      `json:"expectedSalaryEnable" gorm:"column:expectedsalaryenable"`
	ChannelSettingTypeID int64     `json:"channelSettingTypeID" gorm:"column:channelsettingtypeid"`
	IDFrontURL           string    `json:"iDFrontURL" gorm:"column:idfronturl"`
	IDBackURL            string    `json:"iDBackURL" gorm:"column:idbackurl"`
	PortraitURL          string    `json:"portraitURL" gorm:"column:portraiturl"`
	RewardID             int64     `json:"rewardID" gorm:"column:rewardid"`
	PaymentMethodID      int64     `json:"paymentMethodID" gorm:"column:paymentmethodid"`
	TestimonialsID       int64     `json:"testimonialsID" gorm:"column:testimonialsid"`
	VerificationStatus   bool      `json:"verificationStatus" gorm:"column:verificationstatus"`
	Enabled              bool      `json:"enabled" gorm:"column:enabled"`
	ActiveDate           time.Time `json:"activeDate" gorm:"column:activedate"`
	Active               bool      `json:"active" gorm:"column:active"`
	CreatedBy            string    `json:"createdBy" gorm:"column:createdby"`
	CreatedDate          time.Time `json:"createdDate" gorm:"column:createddate"`
	ModifiedBy           string    `json:"modifiedBy" gorm:"column:modifiedby"`
	ModifiedDate         time.Time `json:"modifiedDate" gorm:"column:modifieddate"`
	IsRemove             bool      `json:"isRemove" gorm:"column:isremove"`
	IsOnBoarding         bool      `json:"isOnBoarding" gorm:"column:isonboarding"`
	Code                 string    `json:"code" gorm:"column:code"`
	PortraitRightURL     string    `json:"portraitRightURL" gorm:"column:portraitrighturl"`
	PortraitLeftURL      string    `json:"portraitLeftURL" gorm:"column:portraitlefturl"`
	LivenessStatus       bool      `json:"livenessStatus" gorm:"column:livenessstatus"`
}

func (Kol) TableName() string {
	return Const.TABLE_KOL
}

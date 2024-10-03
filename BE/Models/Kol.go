package Models

import (
	"math/rand"
	"time"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Utils"
)

type Kol struct {
	KolID                int64     `json:"kolID"  gorm:"primaryKey;  column:KolID"`
	UserProfileID        int64     `json:"userProfileID" gorm:"column:UserProfileID"`
	Language             string    `json:"language" gorm:"column:Language"`
	Education            string    `json:"education" gorm:"column:Education"`
	ExpectedSalary       int64     `json:"expectedSalary" gorm:"column:ExpectedSalary"`
	ExpectedSalaryEnable bool      `json:"expectedSalaryEnable" gorm:"column:ExpectedSalaryEnable"`
	ChannelSettingTypeID int64     `json:"channelSettingTypeID" gorm:"column:ChannelSettingTypeID"`
	IDFrontURL           string    `json:"iDFrontURL" gorm:"column:IDFrontURL"`
	IDBackURL            string    `json:"iDBackURL" gorm:"column:IDBackURL"`
	PortraitURL          string    `json:"portraitURL" gorm:"column:PortraitURL"`
	RewardID             int64     `json:"rewardID" gorm:"column:RewardID"`
	PaymentMethodID      int64     `json:"paymentMethodID" gorm:"column:PaymentMethodID"`
	TestimonialsID       int64     `json:"testimonialsID" gorm:"column:TestimonialsID"`
	VerificationStatus   bool      `json:"verificationStatus" gorm:"column:VerificationStatus"`
	Enabled              bool      `json:"enabled" gorm:"column:Enabled"`
	ActiveDate           time.Time `json:"activeDate" gorm:"column:ActiveDate"`
	Active               bool      `json:"active" gorm:"column:Active"`
	CreatedBy            string    `json:"createdBy" gorm:"column:CreatedBy"`
	CreatedDate          time.Time `json:"createdDate" gorm:"column:CreatedDate"`
	ModifiedBy           string    `json:"modifiedBy" gorm:"column:ModifiedBy"`
	ModifiedDate         time.Time `json:"modifiedDate" gorm:"column:ModifiedDate"`
	IsRemove             bool      `json:"isRemove" gorm:"column:IsRemove"`
	IsOnBoarding         bool      `json:"isOnBoarding" gorm:"column:IsOnBoarding"`
	Code                 string    `json:"code" gorm:"column:Code"`
	PortraitRightURL     string    `json:"portraitRightURL" gorm:"column:PortraitRightURL"`
	PortraitLeftURL      string    `json:"portraitLeftURL" gorm:"column:PortraitLeftURL"`
	LivenessStatus       bool      `json:"livenessStatus" gorm:"column:LivenessStatus"`
}

func (Kol) TableName() string {
	return Const.TABLE_KOL
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomKol(i int64) Kol {
	return Kol{
		UserProfileID:        int64(i + 1),
		Language:             randomString(2),
		Education:            randomString(10),
		ExpectedSalary:       rand.Int63n(10000) + 3000,
		ExpectedSalaryEnable: rand.Intn(2) == 1,
		ChannelSettingTypeID: int64(rand.Intn(5) + 1),
		IDFrontURL:           "https://i0.wp.com/sbcf.fr/wp-content/uploads/2018/03/sbcf-default-avatar.png?ssl=1",
		IDBackURL:            "https://i0.wp.com/sbcf.fr/wp-content/uploads/2018/03/sbcf-default-avatar.png?ssl=1",
		PortraitURL:          "https://i0.wp.com/sbcf.fr/wp-content/uploads/2018/03/sbcf-default-avatar.png?ssl=1",
		RewardID:             int64(rand.Intn(999) + 1),
		PaymentMethodID:      int64(rand.Intn(999) + 1),
		TestimonialsID:       int64(rand.Intn(999) + 1),
		VerificationStatus:   rand.Intn(2) == 1,
		Enabled:              rand.Intn(2) == 1,
		ActiveDate:           time.Now(),
		Active:               true,
		CreatedBy:            "admin",
		CreatedDate:          time.Now(),
		ModifiedBy:           "admin",
		ModifiedDate:         time.Now(),
		IsRemove:             false,
		IsOnBoarding:         rand.Intn(2) == 1,
		Code:                 "KOL" + Utils.Int64ToString(rand.Int63n(10000)+3000),
		PortraitRightURL:     "https://i0.wp.com/sbcf.fr/wp-content/uploads/2018/03/sbcf-default-avatar.png?ssl=1",
		PortraitLeftURL:      "https://i0.wp.com/sbcf.fr/wp-content/uploads/2018/03/sbcf-default-avatar.png?ssl=1",
		LivenessStatus:       rand.Intn(2) == 1,
	}
}

package Initializers

import (
	"log"
	"math/rand"
	"time"

	"wan-api-kol-event/Models"
)

func CreateDummyData() {
	// if err := DB.AutoMigrate(&Models.Kol{}); err != nil {
	// 	log.Println("Error during migration:", err)
	// 	return
	// }
	//Đếm dòng dữ liệu
	var count int64
	if err := DB.Model(&Models.Kol{}).Count(&count).Error; err != nil {
		log.Println("Error counting records in Kol table:", err)
		return
	}
	// nếu không đủ dòng dự liệu, tạo data
	if count < 25 {
		rowsNeeded := 25 - count
		dummyData := generateDummyData(int(rowsNeeded))

		if err := DB.Create(&dummyData).Error; err != nil {
			log.Println("Error inserting dummy data:", err)
		} else {
			log.Printf("Inserted %d rows of dummy data to reach 25 total.\n", rowsNeeded)
		}
	} else {
		log.Println("Kol table already has 25 or more rows. Skipping dummy data insertion.")
	}
}
func generateDummyData(num int) []Models.Kol {
	var dummyData []Models.Kol
	for i := 1; i <= num; i++ {
		kol := Models.Kol{
			KolID:                int64(i),
			UserProfileID:        int64(i),
			Language:             randomLanguage(),
			Education:            randomEducation(),
			ExpectedSalary:       rand.Int63n(100000) + 3000,
			ExpectedSalaryEnable: rand.Intn(2) == 1,
			ChannelSettingTypeID: int64(rand.Intn(5) + 1),
			IDFrontURL:           "http://example.com/idfront/" + randomString(3) + ".jpg",
			IDBackURL:            "http://example.com/idback/" + randomString(3) + ".jpg",
			PortraitURL:          "http://example.com/portrait/" + randomString(3) + ".jpg",
			RewardID:             int64(rand.Intn(5) + 1),
			PaymentMethodID:      int64(rand.Intn(5) + 1),
			TestimonialsID:       int64(rand.Intn(5) + 1),
			VerificationStatus:   rand.Intn(2) == 1,
			Enabled:              rand.Intn(2) == 1,
			ActiveDate:           time.Now().AddDate(0, -rand.Intn(6), -rand.Intn(30)),
			Active:               rand.Intn(2) == 1,
			CreatedBy:            "admin",
			CreatedDate:          time.Now().AddDate(0, -rand.Intn(6), -rand.Intn(30)),
			ModifiedBy:           "admin",
			ModifiedDate:         time.Now(),
			IsRemove:             false,
			IsOnBoarding:         rand.Intn(2) == 1,
			Code:                 "KOL" + randomString(3),
			PortraitRightURL:     "http://example.com/portraitright/" + randomString(3) + ".jpg",
			PortraitLeftURL:      "http://example.com/portraitleft/" + randomString(3) + ".jpg",
			LivenessStatus:       rand.Intn(2) == 1,
		}
		dummyData = append(dummyData, kol)
	}
	return dummyData
}
func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func randomLanguage() string {
	languages := []string{"English", "Spanish", "Chinese", "French", "German"}
	return languages[rand.Intn(len(languages))]
}

func randomEducation() string {
	educations := []string{"High School", "Bachelor's Degree", "Master's Degree", "Ph.D.", "Associate Degree"}
	return educations[rand.Intn(len(educations))]
}

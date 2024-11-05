package Initializers

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Models"
	"wan-api-kol-event/ViewModels"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	//Get database url from environment variables (defined in .env file)
	var dsn string = os.Getenv("DB_URL")

	//Connect with postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             50 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Warn,           // Log level
				IgnoreRecordNotFoundError: false,                 // Dont ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,                 // Include params in the SQL log
				Colorful:                  true,                  // Disable color
			},
		),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)

	} else {
		OutputViewModelData()
	}

}

func TransformDataToDTO() ([]*DTO.KolDTO, error) {
	var kols []Models.Kol

	if err := DB.Find(&kols).Error; err != nil {
		log.Println("Error fetching data: ", err)
		return nil, err
	}

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
func OutputViewModelData() {
	kolDTOs, err := TransformDataToDTO()
	if err != nil {
		log.Println("Error transform data to DTO:", err)
		return
	}

	// Create the ViewModel
	response := ViewModels.KolViewModel{
		Result:       "Success",
		ErrorMessage: "",
		PageIndex:    1,
		PageSize:     int64(len(kolDTOs)),
		Guid:         "",
		TotalCount:   int64(len(kolDTOs)),
		KOL:          kolDTOs,
	}

	// Convert response to JSON
	jsonResponse, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Println("Error marshalling response to JSON:", err)
		return
	}

	// Log the JSON response
	log.Println(string(jsonResponse))
}

package Initializers

import (
	"log"
	"os"
	"time"
	"wan-api-kol-event/Models"

	"math/rand"

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
		log.Fatal("Failed to connect to database")
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func tableExists(db *gorm.DB, tableName string) bool {
	var count int64
	db.Table("pg_tables").Where("tablename = ?", tableName).Count(&count)
	return count > 0
}

func MigrateAndSeed() {

	if !tableExists(DB, "KOL") {
		if err := DB.AutoMigrate(&Models.Kol{}); err != nil {
			log.Fatal("Migration failed:", err)
		}
		for i := 0; i < 50; i++ {
			kol := Models.Kol{
				UserProfileID:        int64(i + 1),
				Language:             randomString(2),
				Education:            randomString(10),
				ExpectedSalary:       rand.Int63n(10000) + 3000,
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
				ActiveDate:           time.Now(),
				Active:               true,
				CreatedBy:            "admin",
				CreatedDate:          time.Now(),
				ModifiedBy:           "admin",
				ModifiedDate:         time.Now(),
				IsRemove:             false,
				IsOnBoarding:         rand.Intn(2) == 1,
				Code:                 "KOL" + randomString(3),
				PortraitRightURL:     "http://example.com/portraitright/" + randomString(3) + ".jpg",
				PortraitLeftURL:      "http://example.com/portraitleft/" + randomString(3) + ".jpg",
				LivenessStatus:       rand.Intn(2) == 1,
			}

			if err := DB.Create(&kol).Error; err != nil {
				log.Println("Error inserting fake data:", err)
			}
		}
		log.Println("Migration and seeding completed successfully!")

	} else {
		log.Println("Table 'KOL' already exists. Skipping migration.")
	}
}

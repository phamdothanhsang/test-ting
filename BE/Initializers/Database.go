package Initializers

import (
	"log"
	"os"
	"time"
	"wan-api-kol-event/Models"

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
			kol := Models.GetRandomKol(int64(i))

			if err := DB.Create(&kol).Error; err != nil {
				log.Println("Error inserting fake data:", err)
			}
		}
		log.Println("Migration and seeding completed successfully!")

	} else {
		log.Println("Table 'KOL' already exists. Skipping migration.")
	}
}

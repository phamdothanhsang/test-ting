package Initializers

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global variable to hold the database connection
var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	log.Println("Attempting to connect to database with URL:", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // Output logs to stdout
			logger.Config{
				SlowThreshold:             50 * time.Millisecond, // Log queries slower than 50ms
				LogLevel:                  logger.Warn,           // Log only warnings and errors
				IgnoreRecordNotFoundError: false,                 // Log record not found errors
				ParameterizedQueries:      false,                 // Include query parameters in logs
				Colorful:                  true,                  // Enable colorful output
			},
		),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")
}

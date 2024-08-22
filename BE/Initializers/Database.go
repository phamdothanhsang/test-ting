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

// check table
func tableExists(db *gorm.DB, tableName string) bool {
	var count int
	query := `SELECT count(*) FROM information_schema.tables WHERE table_name = ?`
	db.Raw(query, tableName).Scan(&count)
	return count > 0
}

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

	// Check if table exists
	if !tableExists(DB, "KOL") {
		err = DB.AutoMigrate(&Models.Kol{})
		if err != nil {
			log.Fatal("Failed to migrate database schema:", err)
		}
		log.Println("Table created successfully.")
	} else {
		log.Println("Table already exists, skipping migration.")
	}
}

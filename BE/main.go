package main

import (
	"log"
	"wan-api-kol-event/Controllers"
	"wan-api-kol-event/Initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
	Initializers.MigrateAndSeed() // Migrate random fake data
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Allow all origins
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}))

	// Define your Gin routes here
	r.GET("/kols", Controllers.GetKolsController)

	// Run Gin server
	if err := r.Run(":8081"); err != nil {
		log.Println("Failed to start server")
	}
}

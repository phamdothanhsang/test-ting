package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
}

func main() {
	Initializers.DB.AutoMigrate(&Models.Kol{})
	initDatabaseWithDummiesData("Migrate/dummiesData.json")
}

func getJsonData(filePath string) []Models.Kol {
	var kols []Models.Kol
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicln("Fail to open json file", err)
	}
	jsonErr := json.Unmarshal(fileBytes, &kols)
	if jsonErr != nil {
		log.Panicln("Fail to parse json to KOL", jsonErr)
	}
	return kols
}

func initDatabaseWithDummiesData(filePath string) {
	var count int64
	
	Initializers.DB.Model(&Models.Kol{}).Count(&count)
	fmt.Println("Num Database: ", count)
	if count == 0 {
		kols := getJsonData(filePath)
		Initializers.DB.Create(kols)
	}
}


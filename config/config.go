package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port string
	MangaDexAPIURL string
)

func LoadConfig(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	MangaDexAPIURL = os.Getenv("MANGA_API_URL")
}

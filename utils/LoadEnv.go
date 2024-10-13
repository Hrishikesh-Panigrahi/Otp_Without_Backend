package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	fmt.Printf("\n Loaded the .env file \n")
	return os.Getenv(key)
}

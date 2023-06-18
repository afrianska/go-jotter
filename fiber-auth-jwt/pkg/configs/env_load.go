// ./env_load.go
package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// EnvConf func to get env value
func EnvConf(key string) string {
	// Check load .env file
	err := godotenv.Load(".env");
	if err != nil {
		fmt.Print("Error loading .env file") 
	}

	return os.Getenv(key)
}
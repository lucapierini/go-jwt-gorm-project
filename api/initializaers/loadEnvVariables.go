package initializaers

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"path/filepath"
)

func LoadEnvVariables(){
	// err := godotenv.Load("/Users/danus/RepositoriosGitH_LucaP/go-jwt-gorm-project/.env")
	err := godotenv.Load(filepath.Join(os.Getenv("PWD"), "..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
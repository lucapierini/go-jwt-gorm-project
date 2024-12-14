package initializaers

import (
	// "fmt"
	"os"

	// "path/filepath"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb(){
	var err error
	// posible error en el host
	// dsn = data source name
	dsn := os.Getenv("DB")
	// dsn := godotenv.Load(filepath.Join(os.Getenv("PWD"), "..", ".env"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// fmt.Println("ACAAAAAAAAA" + dsn)
	if err != nil {
		panic("failed to connect database")
	}

}
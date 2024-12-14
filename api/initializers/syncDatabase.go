package initializaers

import (
	"github.com/lucapierini/api/models"
)

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})

	
}
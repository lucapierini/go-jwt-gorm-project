package initializaers

import (
	"log"
	"github.com/joho/godotenv"
	"path/filepath"
)

func LoadEnvVariables(){
	    // Define la ruta al archivo .env en la carpeta padre
		envPath := filepath.Join("..", ".env")
    
		// Carga el archivo .env
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatalf("Error cargando el archivo .env: %v", err)
		}
}
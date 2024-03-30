package configs

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    "banking/models" // Asegúrate de importar tus modelos aquí
)

var DB *gorm.DB

func init() {
	// Cargar el archivo .env
	err := godotenv.Load() // Carga las variables de entorno desde el archivo .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Database connection successfully established")
	DB = db

    
	err = db.AutoMigrate(&models.User{},&models.Banco{} )

    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}

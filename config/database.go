package configs

import (
	"banking/models" // Asegúrate de importar tus modelos aquí
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Cargar el archivo .env
	fmt.Println("Cargando configuracion ENV")
	fmt.Println(os.Getenv("DB_HOST"))
}

func ConnectToDB() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("BANKING_DB")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Database connection successfully established")
	DB = db

	err = db.AutoMigrate(&models.User{}, &models.Banco{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

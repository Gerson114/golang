package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"GO/schemas"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado, usando variáveis do sistema")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no PostgreSQL: %v", err)
	}

	// Migra todas as tabelas para o mesmo banco
	err = db.AutoMigrate(&schemas.User{}, &schemas.Opening{})
	if err != nil {
		log.Fatalf("Erro ao migrar schema: %v", err)
	}

	DB = db
	log.Println("PostgreSQL conectado - Todas as tabelas migradas com sucesso")
}

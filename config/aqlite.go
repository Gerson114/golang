package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"GO/schemas" // <-- importe o model corretamente
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado, usando variáveis do sistema")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no Postgres: %v", err)
	}

	// 🔥 Migrate: cria a tabela openings
	err = db.AutoMigrate(&schemas.Opening{}, &schemas.Create{})
	if err != nil {
		log.Fatalf("Erro ao migrar schema: %v", err)
	}

	DB = db
	log.Println("Banco Postgres conectado e migrado com sucesso")
}

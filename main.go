package main

import (
	"GO/config"
	"GO/router"

	"github.com/joho/godotenv"
)

func main() {
	// Inicializa configurações, banco, etc.
	if err := config.Init(); err != nil {
		// Como logger pode não estar pronto, só printa direto
		println("Erro na configuração:", err.Error())
		return
	}

	logger := config.NewLogger("main ") // cria logger com prefixo "main "

	logger.Infof("Aplicação iniciada com sucesso.")

	godotenv.Load()

	router.Initialize()
}

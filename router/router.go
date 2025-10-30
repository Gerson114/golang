package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Configura modo release para produção
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Desabilita proxy trust para resolver problemas de proxy
	router.SetTrustedProxies(nil)

	InitializeRoutes(router)

	// Usa porta do ambiente ou 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	router.Run(port)
}

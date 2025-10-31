package clientservice

import (
	"net/http"

	"GO/config"
	"GO/schemas"

	"github.com/gin-gonic/gin"
)

func ShowOpenHandler(ctx *gin.Context) {
	// Busca todas as vagas
	var openings []schemas.Opening
	if err := config.DB.
		Preload("User"). // carrega os dados do usuário relacionado
		Find(&openings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados: " + err.Error()})
		return
	}

	// Retorna todas as vagas
	ctx.JSON(http.StatusOK, gin.H{
		"data": openings,
	})
}

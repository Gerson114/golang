package handler

import (
	"net/http"

	"GO/config"
	"GO/schemas"

	"github.com/gin-gonic/gin"
)

func ShowOpenHandler(ctx *gin.Context) {
	// 1. Recupera userId do contexto (injetado pelo middleware JWT)
	userIDValue, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userID, ok := userIDValue.(uint64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Tipo de userId inválido"})
		return
	}

	// 2. Busca vagas apenas do usuário autenticado
	var openings []schemas.Opening
	if err := config.DB.
		Where("user_id = ?", uint(userID)).
		Preload("User").
		Find(&openings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados: " + err.Error()})
		return
	}

	// 3. Retorna os dados
	ctx.JSON(http.StatusOK, gin.H{
		"data": openings,
	})
}

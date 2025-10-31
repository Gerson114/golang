package handler

import (
	"net/http"

	"GO/config"
	"GO/schemas"

	"github.com/gin-gonic/gin"
)

func CreateOpenHandler(ctx *gin.Context) {
	var opening schemas.Opening

	// Faz o bind do corpo JSON
	if err := ctx.ShouldBindJSON(&opening); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Pega o userId do contexto (injetado pelo middleware do JWT)
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

	// Verifica se userID é válido
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserId inválido"})
		return
	}

	// Verifica se o usuário existe
	var user schemas.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Atribui o userId ao registro
	opening.UserId = uint(userID)

	// Cria a vaga no banco
	if err := config.DB.Create(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco: " + err.Error()})
		return
	}

	// Retorna sucesso (sem tentar preload, se não quiser)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registro criado com sucesso",
		"data":    opening,
	})
}

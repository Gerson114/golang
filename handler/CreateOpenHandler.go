package handler

import (
	"net/http"

	"GO/config"
	"GO/schemas"

	"github.com/gin-gonic/gin"
)

func CreateOpenHandler(ctx *gin.Context) {
	var input schemas.Opening

	// Faz o bind do JSON recebido
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Pega o userId do contexto (injetado pelo middleware JWT)
	userIDValue, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userID, ok := userIDValue.(uint64)
	if !ok || userID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserId inválido"})
		return
	}

	// Verifica se o usuário existe
	var user schemas.User
	if err := config.DB.First(&user, uint(userID)).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Cria o registro da vaga, preenchendo apenas os campos existentes
	newOpening := schemas.Opening{
		Role:     input.Role,
		Company:  input.Company,
		Location: input.Location,
		Remote:   input.Remote,
		Link:     input.Link,
		Salary:   input.Salary,
		UserId:   uint(userID),
	}

	// Salva no banco
	if err := config.DB.Create(&newOpening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco: " + err.Error()})
		return
	}

	// Retorna sucesso
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registro criado com sucesso",
		"data":    newOpening,
	})
}

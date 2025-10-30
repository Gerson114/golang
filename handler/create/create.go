package create

import (
	"GO/config"
	"GO/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser é o handler para criar um novo usuário
func CreateUser(c *gin.Context) {
	var input schemas.Create

	// Faz o bind do JSON recebido para a struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PassWord), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criptografar a senha"})
		return
	}

	// Atualiza a senha com o hash
	input.PassWord = string(hashedPassword)

	input.Role = "usuario"

	// Salva o usuário no banco
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco: " + err.Error()})
		return
	}

	// Retorna sucesso
	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

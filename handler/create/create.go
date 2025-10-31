package create

import (
	"GO/config"
	"GO/schemas"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var input schemas.User

	// Bind JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("DEBUG: erro no bind do JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Trim de espaços
	input.PassWord = strings.TrimSpace(input.PassWord)
	input.Email = strings.TrimSpace(input.Email)
	input.Nome = strings.TrimSpace(input.Nome)

	// Gerar hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PassWord), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("DEBUG: erro ao gerar hash da senha:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criptografar a senha"})
		return
	}
	input.PassWord = string(hashedPassword)
	input.Role = "empresa" // padrão

	// Salvar no banco
	if err := config.DB.Create(&input).Error; err != nil {
		fmt.Println("DEBUG: erro ao salvar usuário no banco:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco: " + err.Error()})
		return
	}

	fmt.Println("DEBUG: usuário criado com sucesso - Email:", input.Email)
	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

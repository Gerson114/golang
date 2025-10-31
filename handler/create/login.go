package create

import (
	"GO/config"
	"GO/schemas"
	"GO/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Estrutura de entrada de login

func LoginUser(c *gin.Context) {
	var input schemas.User

	// Bind do JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("DEBUG: erro no bind do JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos. Verifique email e senha."})
		return
	}

	// Trim para evitar espaços extras
	input.Email = strings.TrimSpace(input.Email)
	input.PassWord = strings.TrimSpace(input.PassWord)
	fmt.Println("DEBUG: JSON recebido - Email:", input.Email, "Senha:", input.PassWord)

	// Buscar usuário pelo email
	var user schemas.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		fmt.Println("DEBUG: usuário não encontrado para email:", input.Email, "Erro:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}
	fmt.Println("DEBUG: usuário encontrado - ID:", user.ID, "Email:", user.Email)

	// Comparar senha informada com hash do banco
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(input.PassWord)); err != nil {
		fmt.Println("DEBUG: senha incorreta")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerar token JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		fmt.Println("DEBUG: erro ao gerar JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno ao gerar token JWT."})
		return
	}

	fmt.Println("DEBUG: login bem-sucedido, token gerado")
	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
		"role":    user.Role,
	})
}

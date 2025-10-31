package create

import (
	"GO/config"
	"GO/schemas"
	"GO/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Estrutura de entrada de login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Função para autenticar o usuário
func LoginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("DEBUG: erro ao fazer bind do JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos. Verifique email e senha.",
		})
		return
	}

	fmt.Println("DEBUG: JSON recebido - Email:", input.Email, "Senha:", input.Password)

	// Buscar usuário pelo email
	var user schemas.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		fmt.Println("DEBUG: usuário não encontrado para email:", input.Email, "Erro:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	fmt.Println("DEBUG: usuário encontrado - ID:", user.ID, "Email:", user.Email, "Hash senha:", user.PassWord)

	// Verificar se o hash da senha é válido
	if len(user.PassWord) < 20 {
		fmt.Println("DEBUG: hash da senha inválido ou vazio")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Senha no banco não é um hash válido. Verifique o cadastro do usuário.",
		})
		return
	}

	// Comparar senha informada com o hash do banco
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(input.Password)); err != nil {
		fmt.Println("DEBUG: bcrypt falhou - senha incorreta")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	fmt.Println("DEBUG: senha correta, gerando token JWT")

	// Gerar token JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		fmt.Println("DEBUG: erro ao gerar JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno ao gerar token JWT.",
		})
		return
	}

	fmt.Println("DEBUG: login bem-sucedido, token gerado")

	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
		"role":    user.Role,
	})
}

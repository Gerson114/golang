package create

import (
	"GO/config"
	"GO/schemas"
	"GO/utils"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos. Verifique email e senha.",
		})
		return
	}

	// Buscar usuário pelo email
	var user schemas.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// ⚠️ Verifique se user.Password realmente contém o hash bcrypt
	// Exemplo: "$2a$10$z8sJ..."
	if len(user.PassWord) < 20 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Senha no banco não é um hash válido. Verifique o cadastro do usuário.",
		})
		return
	}

	// Comparar senha informada com o hash do banco
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerar token JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno ao gerar token JWT.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
		"role":    user.Role,
	})
}

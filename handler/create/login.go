package create

import (
	"GO/config"
	"GO/schemas"
	"GO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser autentica o usuário e gera um JWT
func LoginUser(c *gin.Context) {
	// 1. Estrutura para receber os dados de login
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// 2. Faz o bind e valida o JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos ou incompletos"})
		return
	}

	// 3. Declara struct de usuário para buscar no DB
	var user schemas.Create

	// 4. Busca o usuário pelo email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// 5. Compara a senha recebida com o hash armazenado
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// 6. Gera token JWT passando user ID e role
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	// 7. Retorna o token para o cliente
	c.JSON(http.StatusOK, gin.H{"token": token})
}

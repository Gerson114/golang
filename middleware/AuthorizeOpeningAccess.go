package middleware

import (
	"GO/config"
	"GO/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthorizeOpeningAccess permite acesso se for dono da vaga ou admin
func AuthorizeOpeningAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Recupera userId e role do contexto
		userIDValue, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
			c.Abort()
			return
		}
		userID, ok := userIDValue.(uint64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Formato de userId inválido"})
			c.Abort()
			return
		}

		roleValue, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role não encontrado no contexto"})
			c.Abort()
			return
		}
		role, ok := roleValue.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Formato de role inválido"})
			c.Abort()
			return
		}

		// Pega o ID da vaga da URL
		openingIDParam := c.Param("id")
		openingID, err := strconv.ParseUint(openingIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID da vaga inválido"})
			c.Abort()
			return
		}

		// Busca a vaga no banco
		var opening schemas.Opening
		if err := config.DB.First(&opening, openingID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Vaga não encontrada"})
			c.Abort()
			return
		}

		// Se não for admin E não for dono, bloqueia
		if role != "admin" && opening.UserId != uint(userID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Você não tem permissão para acessar esta vaga"})
			c.Abort()
			return
		}

		// Tudo certo, continua
		c.Next()
	}
}

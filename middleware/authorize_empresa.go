package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Apenas usuários com role "empresa"
func AuthorizeEmpresa() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role não encontrada"})
			c.Abort()
			return
		}

		role, ok := roleValue.(string)
		if !ok || role != "empresa" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: apenas empresas podem acessar"})
			c.Abort()
			return
		}

		c.Next()
	}
}

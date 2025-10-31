package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeCliente() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pega a chave secreta do .env
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		// Pega o token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Remove "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		// Parse do token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Pega claims do token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims inválidas"})
			c.Abort()
			return
		}

		// Verifica se o role é "cliente"
		role, ok := claims["role"].(string)
		if !ok || role != "cliente" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: apenas clientes podem acessar"})
			c.Abort()
			return
		}

		// Salva role e userId no contexto
		c.Set("role", role)
		c.Set("userId", claims["userId"])

		c.Next()
	}
}

package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware genérico para verificar role do JWT
func AuthorizeRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		if len(jwtSecret) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET não configurado"})
			c.Abort()
			return
		}

		// 1. Pega token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Remove prefixo "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		// 2. Parse do JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// 3. Pega claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims inválidas"})
			c.Abort()
			return
		}

		// 4. Verifica role
		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: apenas " + requiredRole + " podem acessar"})
			c.Abort()
			return
		}

		// 5. Pega user_id e converte para uint64
		userIdFloat, ok := claims["user_id"].(float64) // JWT numeric sempre vem como float64
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id inválido no token"})
			c.Abort()
			return
		}
		userId := uint64(userIdFloat)

		// 6. Salva role e userId no contexto
		c.Set("role", role)
		c.Set("userId", userId)

		c.Next()
	}
}

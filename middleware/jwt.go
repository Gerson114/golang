package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// 1. Pega token do cookie ou header
		cookieToken, err := c.Cookie("jwt")
		if err == nil && cookieToken != "" {
			tokenString = cookieToken
		} else {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
				c.Abort()
				return
			}
			parts := strings.Fields(authHeader)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
				c.Abort()
				return
			}
			tokenString = parts[1]
		}

		// 2. Pega secret do ambiente
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET não configurado"})
			c.Abort()
			return
		}

		// 3. Parse do token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// 4. Extrai claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims inválidos"})
			c.Abort()
			return
		}

		// 5. Extrai user_id
		var userId uint64
		switch v := claims["user_id"].(type) {
		case float64:
			userId = uint64(v)
		case string:
			userId, err = strconv.ParseUint(v, 10, 64)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id inválido no token"})
				c.Abort()
				return
			}
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id ausente ou inválido no token"})
			c.Abort()
			return
		}

		// 6. Extrai role
		role, ok := claims["role"].(string)
		if !ok || role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role ausente ou inválido no token"})
			c.Abort()
			return
		}

		// 7. Salva userId e role no contexto
		c.Set("userId", userId)
		c.Set("role", role)

		c.Next()
	}
}

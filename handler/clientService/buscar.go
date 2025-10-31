package clientservice

import (
	"net/http"

	"GO/config"
	"GO/schemas"

	"github.com/gin-gonic/gin"
)

func ShowOpenHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	// Busca todas as vagas, carregando o usuário relacionado
	if err := config.DB.Preload("User").Find(&openings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados: " + err.Error()})
		return
	}

	// Monta a resposta com apenas nome e email da empresa
	var result []map[string]interface{}
	for _, o := range openings {
		item := map[string]interface{}{
			"id":       o.ID,
			"role":     o.Role,
			"company":  o.Company,
			"location": o.Location,
			"remote":   o.Remote,
			"link":     o.Link,
			"salary":   o.Salary,
			"user": map[string]interface{}{
				"nome":  o.User.Nome,
				"email": o.User.Email,
			},
		}
		result = append(result, item)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

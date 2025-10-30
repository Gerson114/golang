package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "mensage",
	})
}

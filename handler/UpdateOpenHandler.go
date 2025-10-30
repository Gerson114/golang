package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "mensage",
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "mensage",
	})
}

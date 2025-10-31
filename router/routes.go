package router

import (
	"GO/handler"
	"GO/handler/cliente"
	"GO/handler/create"
	"GO/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	Protect := router.Group("/api/v1")
	Protect.Use(middleware.JWTAuthMiddleware())

	{
		Protect.GET("/buscar", middleware.JWTAuthMiddleware(), handler.ShowOpenHandler)
		Protect.POST("/opening", middleware.JWTAuthMiddleware(), handler.CreateOpenHandler)
		Protect.DELETE("/openingss", handler.DeleteOpenHandler)
		Protect.PUT("/editar", handler.UpdateOpenHandler)
		Protect.GET("/openingr", handler.ListOpenHandler)
	}

	//EMPRESA

	usuario := router.Group("/api/v2")

	usuario.POST("/cadastro", create.CreateUser)
	usuario.POST("/login", create.LoginUser)

	//CLIENTE

	clientes := router.Group("/api/v3")

	clientes.POST("/create", cliente.CreateUser)
	clientes.POST("/login", cliente.LoginUser)

}

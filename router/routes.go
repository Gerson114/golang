package router

import (
	"GO/handler"
	clientservice "GO/handler/clientService"
	"GO/handler/cliente"
	"GO/handler/create"
	"GO/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	// Rotas para empresas
	Protect := router.Group("/api/v1")
	Protect.Use(middleware.AuthorizeRole("empresa")) // middleware genérico

	{
		Protect.GET("/buscar", handler.ShowOpenHandler)
		Protect.POST("/opening", handler.CreateOpenHandler)
		Protect.DELETE("/deletar", handler.DeleteOpenHandler)
		Protect.PUT("/editar", handler.UpdateOpenHandler)
		Protect.GET("/openingr", handler.ListOpenHandler)
	}

	// Rotas para cadastro/login de usuários
	usuario := router.Group("/api/v2")
	{
		usuario.POST("/cadastro", create.CreateUser)
		usuario.POST("/login", create.LoginUser)
	}

	// Rotas para clientes
	clientes := router.Group("/api/v3")
	clientes.Use(middleware.AuthorizeRole("cliente")) // middleware genérico

	{
		clientes.GET("/buscar", clientservice.ShowOpenHandler)
	}

	// Rotas de cadastro/login de clientes
	cli := router.Group("/api/v4")
	{
		cli.POST("/cadastro", cliente.CreateUser)
		cli.POST("/login", cliente.LoginUser)
	}
}

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

	Protect := router.Group("/api/v1")
	Protect.Use(middleware.AuthorizeEmpresa())

	//EMPRESA

	{
		Protect.GET("/buscar", middleware.AuthorizeEmpresa(), handler.ShowOpenHandler)
		Protect.POST("/opening", middleware.AuthorizeEmpresa(), handler.CreateOpenHandler)
		Protect.DELETE("/deletar", middleware.AuthorizeEmpresa(), handler.DeleteOpenHandler)
		Protect.PUT("/editar", handler.UpdateOpenHandler)
		Protect.GET("/openingr", handler.ListOpenHandler)
	}

	usuario := router.Group("/api/v2")

	usuario.POST("/cadastro", create.CreateUser)
	usuario.POST("/login", create.LoginUser)

	//CLIENTE

	clientes := router.Group("/api/v3")
	clientes.Use(middleware.AuthorizeCliente())

	clientes.GET("/buscar", middleware.AuthorizeCliente(), clientservice.ShowOpenHandler)

	cli := router.Group("/api/v4")
	{
		cli.POST("/cadastro", cliente.CreateUser)
		cli.POST("/login", cliente.LoginUser)
	}

}

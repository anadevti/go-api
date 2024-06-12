package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Estabelecendo conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	// Inicializando o repositório de Product
	productRepository := repository.NewProductRepository(dbConnection)

	// Inicializando o usecase de Product
	productUsecase := usecase.NewProductUseCase(productRepository)

	// Inicializando o controlador de Product
	productController := controller.NewProductController(productUsecase)

	// Definindo rotas
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)

	// Iniciando o servidor na porta 8000
	server.Run(":8000")
}

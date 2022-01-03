package main

import (
	"database/sql"

	"github.com/LucasRosello/golang-mvc-template/cmd/server/handler"
	"github.com/LucasRosello/golang-mvc-template/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {

	var db *sql.DB //DB VACIA
	router := gin.Default()

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProduct(productService)
	productRoutes := router.Group("/api/v1/productos")
	{
		productRoutes.GET("/", productHandler.GetAll())
		productRoutes.GET("/:id", productHandler.Get())
		productRoutes.POST("/", productHandler.Store())
		productRoutes.PATCH("/:id", productHandler.Update())
		productRoutes.DELETE("/:id", productHandler.Delete())
	}

	router.Run()
}

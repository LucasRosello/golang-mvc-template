package main

import (
	"database/sql"

	"github.com/LucasRosello/golang-mvc-template/cmd/server/handler"
	"github.com/LucasRosello/golang-mvc-template/internal/example"
	"github.com/gin-gonic/gin"
)

func main() {

	var db *sql.DB //DB VACIA
	router := gin.Default()

	exampleRepository := example.NewRepository(db)
	exampleService := example.NewService(exampleRepository)
	exampleHandler := handler.NewExample(exampleService)
	examplesRoutes := router.Group("/api/v1/examples")
	{
		examplesRoutes.GET("/", exampleHandler.GetAll())
		examplesRoutes.GET("/:id", exampleHandler.Get())
		examplesRoutes.POST("/", exampleHandler.Store())
		examplesRoutes.PATCH("/:id", exampleHandler.Update())
		examplesRoutes.DELETE("/:id", exampleHandler.Delete())
	}

	router.Run()
}

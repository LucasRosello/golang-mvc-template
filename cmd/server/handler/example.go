package handler

import (
	"context"
	"strconv"

	"github.com/LucasRosello/golang-mvc-template/internal/domain"
	"github.com/LucasRosello/golang-mvc-template/internal/example"
	"github.com/gin-gonic/gin"
)

type Example struct {
	exampleService example.Service
}

func NewExample(e example.Service) *Example {
	return &Example{
		exampleService: e,
	}
}

func (e *Example) Get() gin.HandlerFunc {
	type response struct {
		Data domain.Example `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		sel, err := e.exampleService.Get(ctx, int(0))
		if err != nil {
			//c.JSON(404, web.NewError(404, "example not found"))
			c.JSON(404, "Ejemplo no encontrado (o no se implemento el metodo, revisar codigo)")
			return
		}
		c.JSON(201, &response{sel})
	}
}

func (e *Example) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Example `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		examples, err := e.exampleService.GetAll(ctx)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(201, &response{examples})
	}
}

func (e *Example) Store() gin.HandlerFunc {
	type request struct {
		ExampleText string `json:"example-text"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		err := c.Bind(&req)
		if err != nil {
			c.JSON(422, "json decoding: "+err.Error())
			return
		}

		ctx := context.Background()
		example, _ := e.exampleService.Store(ctx, req.ExampleText)

		c.JSON(201, &response{example})
	}
}

func (e *Example) Update() gin.HandlerFunc {
	type request struct {
		ExampleText string `json:"example-text"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		//paramID := c.Param("id")

		if err := c.Bind(&req); err != nil {
			c.JSON(422, "json decoding: "+err.Error())
			return
		}

		ctx := context.Background()
		example, err := e.exampleService.Update(ctx, req.ExampleText)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, &response{example})
	}
}

func (e *Example) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, "invalid ID")
			return
		}

		ctx := context.Background()
		err = e.exampleService.Delete(ctx, int(id))
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, "The example has been deleted")
	}
}

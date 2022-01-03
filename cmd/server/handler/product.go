package handler

import (
	"context"
	"strconv"

	"github.com/LucasRosello/golang-mvc-template/internal/domain"
	"github.com/LucasRosello/golang-mvc-template/internal/product"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(e product.Service) *Product {
	return &Product{
		productService: e,
	}
}

func (e *Product) Get() gin.HandlerFunc {
	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		sel, err := e.productService.Get(ctx, int(0))
		if err != nil {
			//c.JSON(404, web.NewError(404, "product not found"))
			c.JSON(404, "Producto no encontrado (o no se implemento el metodo, revisar codigo)")
			return
		}
		c.JSON(201, &response{sel})
	}
}

func (e *Product) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		products, err := e.productService.GetAll(ctx)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(201, &response{products})
	}
}

func (e *Product) Store() gin.HandlerFunc {
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
		product, _ := e.productService.Store(ctx, req.ExampleText)

		c.JSON(201, &response{product})
	}
}

func (e *Product) Update() gin.HandlerFunc {
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
		product, err := e.productService.Update(ctx, req.ExampleText)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, &response{product})
	}
}

func (e *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, "invalid ID")
			return
		}

		ctx := context.Background()
		err = e.productService.Delete(ctx, int(id))
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, "El producto fue eliminado")
	}
}

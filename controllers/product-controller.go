package controllers

import (
	"github.com/atvd1998/golang-api/entities"
	"github.com/atvd1998/golang-api/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll() []entities.Product
	Save(ctx *gin.Context) entities.Product
	Update(ctx *gin.Context) entities.Product
	Delete(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(ProductService services.ProductService) ProductController {
	return &productController{
		productService: ProductService,
	}
}

func (controller *productController) FindAll() []entities.Product {
	return controller.productService.FindAll()
}

func (controller *productController) Save(ctx *gin.Context) entities.Product {
	var product entities.Product
	ctx.BindJSON(&product)
	controller.productService.Save(product)
	return product
}

func (controller *productController) Update(ctx *gin.Context) entities.Product {
	var product entities.Product
	id := ctx.Params.ByName("id")
	ctx.BindJSON(&product)
	controller.productService.Update(id , product)
	return product
}

func (controller *productController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	controller.productService.Delete(id)
}
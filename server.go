package main

import (
	"github.com/atvd1998/golang-api/controllers"
	"github.com/atvd1998/golang-api/database"
	"github.com/atvd1998/golang-api/repositories"
	"github.com/atvd1998/golang-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	dsn := "host=localhost user=postgres password=tunglatao809 dbname=test port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var (
		productRepository repositories.ProductRepository = repositories.NewProductRepository(db)
		productService    services.ProductService        = services.NewProductService(productRepository)
		productController controllers.ProductController  = controllers.NewProductController(productService)
		// videoService    services.VideoService       = services.New()
		// videoController controllers.VideoController = controllers.New(videoService)
	)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&database.ProductDB{})
	db.Create(&database.ProductDB{Title: "title1", Description: "description1"})
	db.Create(&database.ProductDB{Title: "title2", Description: "description2"})
	db.Create(&database.ProductDB{Title: "title3", Description: "description3"})
	db.Create(&database.ProductDB{Title: "title4", Description: "description4"})

	var products []database.ProductDB
	db.Find((&products))

	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": productController.FindAll()})
	})

	server.POST("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{"data": productController.Save(ctx)})
	})

	server.PATCH("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{"data": productController.Update(ctx)})
	})

	server.DELETE("/products/:id", func(ctx *gin.Context) {
		productController.Delete(ctx)
		ctx.JSON(http.StatusCreated, gin.H{"data": "Success"})
	})

	server.Run(":8080")
}

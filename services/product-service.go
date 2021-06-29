package services

import (
	"github.com/atvd1998/golang-api/database"
	"github.com/atvd1998/golang-api/entities"
	"github.com/atvd1998/golang-api/repositories"
)

type ProductService interface {
	FindAll() []entities.Product
	Save(entities.Product) entities.Product
	Update(string, entities.Product) entities.Product
	Delete(string)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(ProductRepository repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: ProductRepository,
	}
}

func (service *productService) Save(product entities.Product) entities.Product {
	productDB := database.ProductDB{Title: product.Title, Description: product.Description}

	service.productRepository.Save(productDB)

	return product
}

func (service *productService) FindAll() []entities.Product {
	productDBs := service.productRepository.FindAll()

	var res []entities.Product
	for _, product := range productDBs {
		res = append(res, entities.Product{Title: product.Title, Description: product.Description})
	}

	return res
}

func (service *productService) Update(id string, product entities.Product) entities.Product {
	productDB := database.ProductDB{Title: product.Title, Description: product.Description}

	service.productRepository.Update(id, productDB)

	return product
}

func (service *productService) Delete(id string) {
	service.productRepository.Delete(id)
}

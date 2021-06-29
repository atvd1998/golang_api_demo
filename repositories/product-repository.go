package repositories

import (
	"github.com/atvd1998/golang-api/database"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []database.ProductDB
	Save(database.ProductDB) database.ProductDB
	Update(string, database.ProductDB) database.ProductDB
	Delete(string) 
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) Save(product database.ProductDB) database.ProductDB {
	repo.db.Create(&product)
	return product
}

func (repo *productRepository) FindAll() []database.ProductDB {
	var products []database.ProductDB
	repo.db.Find(&products)
	return products
}

func (repo *productRepository) Update(id string, product database.ProductDB) database.ProductDB {
	repo.db.Model(&database.ProductDB{}).Where("id = ?", id).Updates(product)
	return product
}

func (repo *productRepository) Delete(id string,) {
	repo.db.Delete(&database.ProductDB{}, id)
}

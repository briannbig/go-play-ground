package service

import (
	"gorm-concept/db"
	"gorm-concept/model"
)

type ProductService interface {
	CreateProduct(name string, price uint) model.Product
	GetAllProducts() []model.Product
	GetProductById(id uint) (model.Product, error)
	UpdateProduct(id uint, name string, price uint)
	DeleteProduct(id uint)
}

type productService struct {
	db db.Database
}

func New() ProductService {
	return productService{
		db: db.New(),
	}
}

func (p productService) CreateProduct(name string, price uint) model.Product {
	id := p.db.Db.Create(&model.Product{Name: name, Price: price})

	var product model.Product

	p.db.Db.First(&product, id)

	return product

}

func (p productService) DeleteProduct(id uint) {
	p.db.Db.Delete(&model.Product{}, id)
}

func (p productService) GetAllProducts() []model.Product {
	var products []model.Product
	p.db.Db.Find(&products)
	return products
}

func (p productService) GetProductById(id uint) (model.Product, error) {
	var product model.Product
	p.db.Db.First(&product, id)
	return product, nil

}

func (p productService) UpdateProduct(id uint, name string, price uint) {
	var product model.Product

	p.db.Db.First(&product, id)

	if name != "" {
		product.Name = name
	}

	if price > 0 {
		product.Price = price
	}

	p.db.Db.Save(&product)
}

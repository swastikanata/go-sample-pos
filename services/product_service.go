package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.Repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.Repo.DeleteProduct(id)
}

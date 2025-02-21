package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, "product_id = ?", id).Error
	return &product, err
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id string) error {
	return r.DB.Delete(&models.Product{}, "product_id = ?", id).Error
}

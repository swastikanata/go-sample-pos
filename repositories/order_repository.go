package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id string) (*models.Order, error)
	Update(id string, order *models.Order) error
	Delete(id string) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (r *orderRepositoryImpl) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepositoryImpl) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepositoryImpl) GetByID(id string) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, "order_id = ?", id).Error
	return &order, err
}

func (r *orderRepositoryImpl) Update(id string, order *models.Order) error {
	return r.db.Model(&models.Order{}).Where("order_id = ?", id).Updates(order).Error
}

func (r *orderRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Order{}, "order_id = ?", id).Error
}

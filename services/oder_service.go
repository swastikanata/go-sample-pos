package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type OrderService interface {
	CreateOrder(product *models.Order) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
	UpdateOrder(id string, product *models.Order) error
	DeleteOrder(id string) error
}

type orderServiceImpl struct {
	Repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderServiceImpl{repo}
}
func (s *orderServiceImpl) CreateOrder(order *models.Order) error {
	return s.Repo.Create(order)
}

func (s *orderServiceImpl) GetAllOrders() ([]models.Order, error) {
	return s.Repo.GetAll()
}

func (s *orderServiceImpl) GetOrderByID(id string) (*models.Order, error) {
	return s.Repo.GetByID(id)
}

func (s *orderServiceImpl) UpdateOrder(id string, order *models.Order) error {
	return s.Repo.Update(id, order)
}

func (s *orderServiceImpl) DeleteOrder(id string) error {
	return s.Repo.Delete(id)
}

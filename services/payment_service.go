package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type PaymentService interface {
	CreatePayment(product *models.Payment) error
	GetAllPayments() ([]models.Payment, error)
	GetPaymentByID(id string) (*models.Payment, error)
	UpdatePayment(id string, product *models.Payment) error
	DeletePayment(id string) error
}

type paymentServiceImpl struct {
	Repo repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository) PaymentService {
	return &paymentServiceImpl{repo}
}
func (s *paymentServiceImpl) CreatePayment(payment *models.Payment) error {
	return s.Repo.Create(payment)
}

func (s *paymentServiceImpl) GetAllPayments() ([]models.Payment, error) {
	return s.Repo.GetAll()
}

func (s *paymentServiceImpl) GetPaymentByID(id string) (*models.Payment, error) {
	return s.Repo.GetByID(id)
}

func (s *paymentServiceImpl) UpdatePayment(id string, payment *models.Payment) error {
	return s.Repo.Update(id, payment)
}

func (s *paymentServiceImpl) DeletePayment(id string) error {
	return s.Repo.Delete(id)
}

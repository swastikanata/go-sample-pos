package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type DiscountService interface {
	CreateDiscount(product *models.Discount) error
	GetAllDiscounts() ([]models.Discount, error)
	GetDiscountByID(id string) (*models.Discount, error)
	UpdateDiscount(id string, product *models.Discount) error
	DeleteDiscount(id string) error
}

type discountServiceImpl struct {
	Repo repositories.DiscountRepository
}

func NewDiscountService(repo repositories.DiscountRepository) DiscountService {
	return &discountServiceImpl{repo}
}
func (s *discountServiceImpl) CreateDiscount(discount *models.Discount) error {
	return s.Repo.Create(discount)
}

func (s *discountServiceImpl) GetAllDiscounts() ([]models.Discount, error) {
	return s.Repo.GetAll()
}

func (s *discountServiceImpl) GetDiscountByID(id string) (*models.Discount, error) {
	return s.Repo.GetByID(id)
}

func (s *discountServiceImpl) UpdateDiscount(id string, discount *models.Discount) error {
	return s.Repo.Update(id, discount)
}

func (s *discountServiceImpl) DeleteDiscount(id string) error {
	return s.Repo.Delete(id)
}

package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type TaxService interface {
	CreateTax(product *models.Tax) error
	GetAllTaxes() ([]models.Tax, error)
	GetTaxByID(id string) (*models.Tax, error)
	UpdateTax(id string, product *models.Tax) error
	DeleteTax(id string) error
}

type taxServiceImpl struct {
	Repo repositories.TaxRepository
}

func NewTaxService(repo repositories.TaxRepository) TaxService {
	return &taxServiceImpl{repo}
}
func (s *taxServiceImpl) CreateTax(tax *models.Tax) error {
	return s.Repo.Create(tax)
}

func (s *taxServiceImpl) GetAllTaxes() ([]models.Tax, error) {
	return s.Repo.GetAll()
}

func (s *taxServiceImpl) GetTaxByID(id string) (*models.Tax, error) {
	return s.Repo.GetByID(id)
}

func (s *taxServiceImpl) UpdateTax(id string, tax *models.Tax) error {
	return s.Repo.Update(id, tax)
}

func (s *taxServiceImpl) DeleteTax(id string) error {
	return s.Repo.Delete(id)
}

package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type TaxRepository interface {
	Create(tax *models.Tax) error
	GetAll() ([]models.Tax, error)
	GetByID(id string) (*models.Tax, error)
	Update(id string, tax *models.Tax) error
	Delete(id string) error
}

type taxRepositoryImpl struct {
	db *gorm.DB
}

func NewTaxRepository(db *gorm.DB) TaxRepository {
	return &taxRepositoryImpl{db}
}

func (r *taxRepositoryImpl) Create(tax *models.Tax) error {
	return r.db.Create(tax).Error
}

func (r *taxRepositoryImpl) GetAll() ([]models.Tax, error) {
	var taxes []models.Tax
	err := r.db.Find(&taxes).Error
	return taxes, err
}

func (r *taxRepositoryImpl) GetByID(id string) (*models.Tax, error) {
	var tax models.Tax
	err := r.db.First(&tax, "tax_id = ?", id).Error
	return &tax, err
}

func (r *taxRepositoryImpl) Update(id string, tax *models.Tax) error {
	return r.db.Model(&models.Tax{}).Where("tax_id = ?", id).Updates(tax).Error
}

func (r *taxRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Tax{}, "tax_id = ?", id).Error
}

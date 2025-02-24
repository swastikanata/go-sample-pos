package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(Discount *models.Discount) error
	GetAll() ([]models.Discount, error)
	GetByID(id string) (*models.Discount, error)
	Update(id string, discount *models.Discount) error
	Delete(id string) error
}

type DiscountRepositoryImpl struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db}
}

func (r *DiscountRepositoryImpl) Create(discount *models.Discount) error {
	return r.db.Create(discount).Error
}

func (r *DiscountRepositoryImpl) GetAll() ([]models.Discount, error) {
	var discounts []models.Discount
	err := r.db.Find(&discounts).Error
	return discounts, err
}

func (r *DiscountRepositoryImpl) GetByID(id string) (*models.Discount, error) {
	var discount models.Discount
	err := r.db.First(&discount, "discount_id = ?", id).Error
	return &discount, err
}

func (r *DiscountRepositoryImpl) Update(id string, discount *models.Discount) error {
	return r.db.Model(&models.Discount{}).Where("discount_id = ?", id).Updates(discount).Error
}

func (r *DiscountRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Discount{}, "discount_id = ?", id).Error
}

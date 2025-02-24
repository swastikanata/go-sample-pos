package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetAll() ([]models.Payment, error)
	GetByID(id string) (*models.Payment, error)
	Update(id string, payment *models.Payment) error
	Delete(id string) error
}

type paymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepositoryImpl{db}
}

func (r *paymentRepositoryImpl) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepositoryImpl) GetAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *paymentRepositoryImpl) GetByID(id string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "payment_id = ?", id).Error
	return &payment, err
}

func (r *paymentRepositoryImpl) Update(id string, payment *models.Payment) error {
	return r.db.Model(&models.Payment{}).Where("payment_id = ?", id).Updates(payment).Error
}

func (r *paymentRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Payment{}, "payment_id = ?", id).Error
}

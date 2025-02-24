package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type ReceiptService interface {
	CreateReceipt(product *models.Receipt) error
	GetAllReceipts() ([]models.Receipt, error)
	GetReceiptByID(id string) (*models.Receipt, error)
	UpdateReceipt(id string, product *models.Receipt) error
	DeleteReceipt(id string) error
}

type receiptServiceImpl struct {
	Repo repositories.ReceiptRepository
}

func NewReceiptService(repo repositories.ReceiptRepository) ReceiptService {
	return &receiptServiceImpl{repo}
}
func (s *receiptServiceImpl) CreateReceipt(receipt *models.Receipt) error {
	return s.Repo.Create(receipt)
}

func (s *receiptServiceImpl) GetAllReceipts() ([]models.Receipt, error) {
	return s.Repo.GetAll()
}

func (s *receiptServiceImpl) GetReceiptByID(id string) (*models.Receipt, error) {
	return s.Repo.GetByID(id)
}

func (s *receiptServiceImpl) UpdateReceipt(id string, receipt *models.Receipt) error {
	return s.Repo.Update(id, receipt)
}

func (s *receiptServiceImpl) DeleteReceipt(id string) error {
	return s.Repo.Delete(id)
}

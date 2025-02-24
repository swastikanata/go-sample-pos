package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type InventoryService interface {
	CreateInventory(product *models.Inventory) error
	GetAllInventories() ([]models.Inventory, error)
	GetInventoryByID(id string) (*models.Inventory, error)
	UpdateInventory(id string, product *models.Inventory) error
	DeleteInventory(id string) error
}

type inventoryServiceImpl struct {
	Repo repositories.InventoryRepository
}

func NewInventoryService(repo repositories.InventoryRepository) InventoryService {
	return &inventoryServiceImpl{repo}
}
func (s *inventoryServiceImpl) CreateInventory(inventory *models.Inventory) error {
	return s.Repo.Create(inventory)
}

func (s *inventoryServiceImpl) GetAllInventories() ([]models.Inventory, error) {
	return s.Repo.GetAll()
}

func (s *inventoryServiceImpl) GetInventoryByID(id string) (*models.Inventory, error) {
	return s.Repo.GetByID(id)
}

func (s *inventoryServiceImpl) UpdateInventory(id string, inventory *models.Inventory) error {
	return s.Repo.Update(id, inventory)
}

func (s *inventoryServiceImpl) DeleteInventory(id string) error {
	return s.Repo.Delete(id)
}

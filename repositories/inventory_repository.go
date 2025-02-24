package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Create(inventory *models.Inventory) error
	GetAll() ([]models.Inventory, error)
	GetByID(id string) (*models.Inventory, error)
	Update(id string, inventory *models.Inventory) error
	Delete(id string) error
}

type inventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepositoryImpl{db}
}

func (r *inventoryRepositoryImpl) Create(inventory *models.Inventory) error {
	return r.db.Create(inventory).Error
}

func (r *inventoryRepositoryImpl) GetAll() ([]models.Inventory, error) {
	var inventories []models.Inventory
	err := r.db.Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepositoryImpl) GetByID(id string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.First(&inventory, "inventory_id = ?", id).Error
	return &inventory, err
}

func (r *inventoryRepositoryImpl) Update(id string, inventory *models.Inventory) error {
	return r.db.Model(&models.Inventory{}).Where("inventory_id = ?", id).Updates(inventory).Error
}

func (r *inventoryRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Inventory{}, "inventory_id = ?", id).Error
}

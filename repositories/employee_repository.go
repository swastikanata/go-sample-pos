package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(customer *models.Employee) error
	GetAll() ([]models.Employee, error)
	GetByID(id string) (*models.Employee, error)
	Update(id string, employee *models.Employee) error
	Delete(id string) error
}

type employeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db}
}

func (r *employeeRepositoryImpl) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepositoryImpl) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepositoryImpl) GetByID(id string) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, "employee_id = ?", id).Error
	return &employee, err
}

func (r *employeeRepositoryImpl) Update(id string, employee *models.Employee) error {
	return r.db.Model(&models.Employee{}).Where("employee_id = ?", id).Updates(employee).Error
}

func (r *employeeRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Employee{}, "employee_id = ?", id).Error
}

package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) error
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	UpdateEmployee(id string, employee *models.Employee) error
	DeleteEmployee(id string) error
}

type employeeSerciveImpl struct {
	Repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeSerciveImpl{repo}
}
func (s *employeeSerciveImpl) CreateEmployee(employee *models.Employee) error {
	return s.Repo.Create(employee)
}

func (s *employeeSerciveImpl) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAll()
}

func (s *employeeSerciveImpl) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.Repo.GetByID(id)
}

func (s *employeeSerciveImpl) UpdateEmployee(id string, employee *models.Employee) error {
	return s.Repo.Update(id, employee)
}

func (s *employeeSerciveImpl) DeleteEmployee(id string) error {
	return s.Repo.Delete(id)
}

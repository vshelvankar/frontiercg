package repository

import (
	models "github.com/vshelvankar/frontiercg/carmgmt/models"
)

// CarsRepository is the repository interface for operations on Cars
type CarsRepository interface {
	GetAll() ([]models.Car, error)
	GetByID(id string) (*models.Car, error)
	Create(car *models.Car) (string, error)
	Delete(id string) error
}

package datastore

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	m "github.com/vshelvankar/frontiercg/carmgmt/models"
	r "github.com/vshelvankar/frontiercg/carmgmt/repository"
)

// CarsDataStore is Inmemory datastore for cars
// To have faster getById there is CarsCacheByID which is like a cache for faster access use case with trade off of redundant car
type CarsDataStore struct {
	Cars          []m.Car
	CarsCacheByID map[string]m.Car
}

// NewCarsInMemoryDataStore is to create CarsDataStore with default values
func NewCarsInMemoryDataStore() r.CarsRepository {
	return &CarsDataStore{
		Cars:          make([]m.Car, 0),
		CarsCacheByID: make(map[string]m.Car),
	}
}

// GetAll is a function to get all cars in datastore
func (cds *CarsDataStore) GetAll() ([]m.Car, error) {
	// Since in memory and store is in our control sending error as nil.
	// If db/other store propagate error
	return cds.Cars, nil
}

// GetByID is a function to get all cars in datastore
func (cds *CarsDataStore) GetByID(id string) (*m.Car, error) {
	// Check if car by id exist in cache. Our cache is always upto date with cars store for easy and fast access
	if car, ok := cds.CarsCacheByID[id]; ok {
		return &car, nil
	}
	return nil, fmt.Errorf("Car by id : %s does not exist", id)
}

// Create is a function to add Car to datastore
func (cds *CarsDataStore) Create(car *m.Car) (string, error) {
	// Add uuid to Car
	id := uuid.NewV4().String()
	car.ID = id
	cds.Cars = append(cds.Cars, *car)
	// Updating cache as well
	cds.CarsCacheByID[car.ID] = *car
	// Since in memory and store is in our control sending error as nil.
	// If db/other store propagate error
	return id, nil
}

// Delete is a function to delete Car to datastore
func (cds *CarsDataStore) Delete(id string) error {
	// Check if cat by ID exist in DS. If not throw error
	if _, ok := cds.CarsCacheByID[id]; !ok {
		return fmt.Errorf("Cannot delete car. Car by id : %s does not exist", id)
	}
	// delete from datastore
	for i, car := range cds.Cars {
		if car.ID == id {
			cds.Cars = append(cds.Cars[:i], cds.Cars[i+1:]...)
			break
		}
	}
	// delete from cache
	delete(cds.CarsCacheByID, id)
	// Since in memory and store is in our control sending error as nil.
	// If db/other store propagate error
	return nil
}

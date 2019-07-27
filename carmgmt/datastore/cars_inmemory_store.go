package datastore

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	m "github.com/vshelvankar/frontiercg/carmgmt/models"
	r "github.com/vshelvankar/frontiercg/carmgmt/repository"
)

// CarsDataStore is Inmemory datastore for cars
// To have faster getById there is carsCacheByID which is like a cache for faster access use case with trade off of extra memory/storage/space
type carsDataStore struct {
	cars []m.Car
	// Cache storing id vs index in store slice
	carsCacheByID map[string]int
}

// NewCarsInMemoryDataStore is to create CarsDataStore with default values
func NewCarsInMemoryDataStore() r.CarsRepository {
	return &carsDataStore{
		cars:          make([]m.Car, 0),
		carsCacheByID: make(map[string]int),
	}
}

// GetAll is a function to get all cars in datastore
func (cds *carsDataStore) GetAll() ([]m.Car, error) {
	// Since in memory and store is in our control sending error as nil.
	// If db/other store propagate error
	return cds.cars, nil
}

// GetByID is a function to get all cars in datastore
func (cds *carsDataStore) GetByID(id string) (*m.Car, error) {
	// Check if car by id exist in cache. Our cache is always upto date with cars store for easy and fast access
	if carIndex, ok := cds.carsCacheByID[id]; ok {
		return &cds.cars[carIndex], nil
	}
	return nil, fmt.Errorf("Car by id : %s does not exist", id)
}

// Create is a function to add Car to datastore
func (cds *carsDataStore) Create(car *m.Car) (string, error) {
	// Add uuid to Car
	id := uuid.NewV4().String()
	car.ID = id
	cds.cars = append(cds.cars, *car)
	// Updating cache as well
	cds.carsCacheByID[car.ID] = len(cds.cars) - 1
	// Since in memory and store is in our control sending error as nil.
	// If db/other store propagate error
	return id, nil
}

// Delete is a function to delete Car to datastore
func (cds *carsDataStore) Delete(id string) error {
	// Check if cat by ID exist in DS. If not throw error
	if _, ok := cds.carsCacheByID[id]; !ok {
		return fmt.Errorf("Cannot delete car. Car by id : %s does not exist", id)
	}
	// delete from datastore
	for i, car := range cds.cars {
		if car.ID == id {
			cds.cars = append(cds.cars[:i], cds.cars[i+1:]...)
			break
		}
	}
	// delete from cache
	delete(cds.carsCacheByID, id)
	// Since in memory and store is in our control sending error as nil.
	// If db/other store propagate error
	return nil
}

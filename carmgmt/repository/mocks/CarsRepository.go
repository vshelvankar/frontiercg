// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/vshelvankar/frontiercg/carmgmt/models"

// CarsRepository is an autogenerated mock type for the CarsRepository type
type CarsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: car
func (_m *CarsRepository) Create(car *models.Car) (string, error) {
	ret := _m.Called(car)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Car) string); ok {
		r0 = rf(car)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Car) error); ok {
		r1 = rf(car)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *CarsRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *CarsRepository) GetAll() ([]models.Car, error) {
	ret := _m.Called()

	var r0 []models.Car
	if rf, ok := ret.Get(0).(func() []models.Car); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Car)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *CarsRepository) GetByID(id string) (*models.Car, error) {
	ret := _m.Called(id)

	var r0 *models.Car
	if rf, ok := ret.Get(0).(func(string) *models.Car); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Car)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

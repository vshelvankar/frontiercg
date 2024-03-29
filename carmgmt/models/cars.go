package models

import (
	"errors"
	"net/http"
	"strconv"

	e "github.com/vshelvankar/frontiercg/carmgmt/errors"
)

// Car struct to hold all car properties
type Car struct {
	ID    string `json:"id,omitempty"`
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
	Year  string `json:"year,omitempty"`
}

// CarResponse struct to hold all car properties for Get Car by id usecase
type CarResponse struct {
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
	Year  string `json:"year,omitempty"`
}

// CarRequest struct to hold all car properties for Get Car by id usecase
type CarRequest struct {
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
	Year  string `json:"year,omitempty"`
}

// NewCarResponse for car details without ID
func NewCarResponse(car *Car) *CarResponse {
	carResponse := &CarResponse{
		Make:  car.Make,
		Model: car.Model,
		Year:  car.Year,
	}
	return carResponse
}

// NewCar from CarRequest
func NewCar(cr *CarRequest) *Car {
	car := &Car{
		Make:  cr.Make,
		Model: cr.Model,
		Year:  cr.Year,
	}
	return car
}

// Bind method to make Car a Render interface. will run after the unmarshalling is complete
func (c *CarRequest) Bind(r *http.Request) error {
	// validations for required fields in Car
	if c.Make == "" {
		return e.New(e.MissingMakeFieldForCreateCode, e.MissingMakeFieldForCreateMsg)
	}

	if c.Model == "" {
		return e.New(e.MissingModelFieldForCreateCode, e.MissingModelFieldForCreateMsg)
	}

	if c.Year == "" {
		return e.New(e.MissingYearFieldForCreateCode, e.MissingYearFieldForCreateMsg)
	}

	if _, err := strconv.Atoi(c.Year); err != nil {
		return errors.New("Year field in not a number in request payload")
	}
	return nil
}

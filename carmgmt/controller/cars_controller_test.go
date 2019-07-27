package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	carerrors "github.com/vshelvankar/frontiercg/carmgmt/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	m "github.com/vshelvankar/frontiercg/carmgmt/models"
	"github.com/vshelvankar/frontiercg/carmgmt/repository/mocks"

	"github.com/go-chi/chi"
)

var c1 = m.Car{ID: "123", Make: "Volvo", Model: "XC60", Year: "1999"}
var c2 = m.Car{ID: "456", Make: "Honda", Model: "Civic", Year: "2015"}
var c3 = m.Car{ID: "789", Make: "Hyndai", Model: "I20", Year: "2000"}

// Create Car Payload
var createCarPayload = `{"make": "Volvo","model": "model-1","year": "2000"}`

// Create Car Payload
var invalidCreateCarPayload = `{"model": "model-1","year": "2000"}`

func TestGetAllCarsIfNothingInRepo(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("GetAll").Return(make([]m.Car, 0), nil)

	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars", cc.GetAllCars)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := make([]m.Car, 0)
	// From response
	actual := make([]m.Car, 0)
	json.Unmarshal(rr.Body.Bytes(), &actual)

	assert.Equal(t, expected, actual)
}

func TestGetAllCars(t *testing.T) {
	// setup expectations from carsRepo
	cars := []m.Car{c1, c2, c3}
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("GetAll").Return(cars, nil)

	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars", cc.GetAllCars)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := cars
	// From response
	actual := make([]m.Car, 0)
	json.Unmarshal(rr.Body.Bytes(), &actual)

	assert.Equal(t, expected, actual)
}

func TestGetAllCarsErrCase(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("GetAll").Return(nil, errors.New("some error"))

	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars", cc.GetAllCars)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestCreateCarSuccess(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carIDCreated := "123"
	carsRepo.On("Create", mock.AnythingOfType("*models.Car")).Return(carIDCreated, nil)

	req, err := http.NewRequest("POST", "/cars", strings.NewReader(createCarPayload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars", cc.CreateCar)

	r.ServeHTTP(rr, req)

	expected := map[string]string{"id": carIDCreated}

	// From response
	var actual map[string]string
	json.Unmarshal(rr.Body.Bytes(), &actual)

	// assert that the expectations were met
	assert.Equal(t, expected, actual)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func TestCreateCarErrCase(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("Create", mock.AnythingOfType("*models.Car")).Return("", errors.New("Some error"))

	req, err := http.NewRequest("POST", "/cars", strings.NewReader(createCarPayload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars", cc.CreateCar)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestCreateCarBindErrCase(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("Create", mock.AnythingOfType("*models.Car")).Return("123", nil)

	req, err := http.NewRequest("POST", "/cars", strings.NewReader(invalidCreateCarPayload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars", cc.CreateCar)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// From response
	var actual map[string]string
	json.Unmarshal(rr.Body.Bytes(), &actual)

	assert.NotNil(t, actual)

}

func TestDeleteCarSuccess(t *testing.T) {
	// setup expectations from carsRepo
	id := "123"
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("Delete", mock.AnythingOfType("string")).Return(nil)

	req, err := http.NewRequest("DELETE", "/cars/"+id, strings.NewReader(createCarPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars/{id}", cc.DeleteCar)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

}

func TestDeleteCarErrNotFoundCase(t *testing.T) {
	// setup expectations from carsRepo
	id := "123"
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("Delete", mock.AnythingOfType("string")).Return(carerrors.CarNotFoundErr(id))

	req, err := http.NewRequest("DELETE", "/cars/"+id, strings.NewReader(createCarPayload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars/{id}", cc.DeleteCar)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}

func TestDeleteCarErrOtherErrCase(t *testing.T) {
	// setup expectations from carsRepo
	id := "123"
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("other error"))

	req, err := http.NewRequest("DELETE", "/cars/"+id, strings.NewReader(createCarPayload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars/{id}", cc.DeleteCar)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestGetCarById(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	carsRepo.On("GetByID", mock.AnythingOfType("string")).Return(&c1, nil)
	req, err := http.NewRequest("GET", "/cars/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars/{id}", cc.GetCarByID)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := map[string]string{"make": "Volvo", "model": "XC60", "year": "1999"}

	// From response
	var actual map[string]string
	json.Unmarshal(rr.Body.Bytes(), &actual)

	// assert that the expectations were met
	assert.Equal(t, expected, actual)
}

func TestGetCarByIdNotFound(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	id := "1234"
	carsRepo.On("GetByID", mock.AnythingOfType("string")).Return(nil, carerrors.CarNotFoundErr(id))
	req, err := http.NewRequest("GET", "/cars/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	//Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars/{id}", cc.GetCarByID)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestGetCarByIdOtherErr(t *testing.T) {
	// setup expectations from carsRepo
	// Creating controller with mocking repository for cars
	var carsRepo = new(mocks.CarsRepository)
	var cc = CarsController{
		repo: carsRepo,
	}
	id := "1234"
	carsRepo.On("GetByID", mock.AnythingOfType("string")).Return(nil, errors.New("Some other error other than CarError. Like DB err"))
	req, err := http.NewRequest("GET", "/cars/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	//Recorder
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/cars/{id}", cc.GetCarByID)

	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

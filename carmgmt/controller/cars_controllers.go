package controller

import (
	"net/http"

	"github.com/vshelvankar/frontiercg/carmgmt/logger"

	m "github.com/vshelvankar/frontiercg/carmgmt/core/models"

	r "github.com/vshelvankar/frontiercg/carmgmt/core/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CarsController - struct for car controller type
type CarsController struct {
	r.CarsRepository
}

// CarHandlerRoutes is a function to configure handlers for car API
func CarHandlerRoutes(repo r.CarsRepository) *chi.Mux {
	carHandler := &CarsController{
		repo,
	}
	router := chi.NewRouter()
	router.Get("/{carID}", carHandler.GetCarByID)
	router.Delete("/{carID}", carHandler.DeleteCar)
	router.Post("/", carHandler.CreateCar)
	router.Get("/", carHandler.GetAllCars)
	logger.Log().Info("Configured Car Handlers")
	return router
}

// GetCarByID is a handler function to get car by ID
func (cc CarsController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "carID")
	car, err := cc.GetByID(carID)
	if err != nil {
		logger.Log().WithError(err).Error("Error during GetCarByID")
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, m.NewCarResponse(car))
}

// DeleteCar is handler function to delete car by ID
func (cc CarsController) DeleteCar(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	carID := chi.URLParam(r, "carID")
	err := cc.Delete(carID)
	if err != nil {
		logger.Log().WithError(err).Error("Error during DeleteCar")
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	response["message"] = "Deleted Car successfully"
	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, response)

}

// CreateCar is handler function to create
func (cc CarsController) CreateCar(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	cr := &m.CarRequest{}
	if err := render.Bind(r, cr); err != nil {
		logger.Log().WithError(err).Error("Error during CreateCar")
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id, err := cc.Create(m.NewCar(cr))
	if err != nil {
		logger.Log().WithError(err).Error("Error during CreateCar")
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	response["id"] = id
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}

// GetAllCars is handler function to get all cars
func (cc CarsController) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := cc.GetAll()
	if err != nil {
		logger.Log().WithError(err).Error("Error during GetAllCars")
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, cars)
}

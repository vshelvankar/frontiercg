package controller

import (
	"net/http"

	carerrors "github.com/vshelvankar/frontiercg/carmgmt/errors"

	"github.com/vshelvankar/frontiercg/carmgmt/logger"

	m "github.com/vshelvankar/frontiercg/carmgmt/models"

	r "github.com/vshelvankar/frontiercg/carmgmt/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CarsController - struct for car controller type
type CarsController struct {
	repo r.CarsRepository
}

// CarHandlerRoutes is a function to configure handlers for car API
func CarHandlerRoutes(repo r.CarsRepository) *chi.Mux {
	carHandler := &CarsController{
		repo: repo,
	}
	router := chi.NewRouter()
	router.Get("/{id}", carHandler.GetCarByID)
	router.Delete("/{id}", carHandler.DeleteCar)
	router.Post("/", carHandler.CreateCar)
	router.Get("/", carHandler.GetAllCars)
	logger.Log().Info("Configured Car Handlers")
	return router
}

// GetCarByID is a handler function to get car by ID
func (cc CarsController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	car, err := cc.repo.GetByID(carID)
	if err != nil {
		logger.Log().WithError(err).Error("Error during GetCarByID")
		switch v := err.(type) {
		case carerrors.CarError:
			// If car not found error.
			if v.Code == carerrors.NotFoundErrCode {
				render.Render(w, r, carerrors.HTTPErrNotFound(err))
				return
			}
		default:
			// If error due to other cases. This will be mostly internal server error
			render.Render(w, r, carerrors.HTTPErrInternalServer(err))
			return
		}
	}
	render.JSON(w, r, m.NewCarResponse(car))
}

// DeleteCar is handler function to delete car by ID
func (cc CarsController) DeleteCar(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	carID := chi.URLParam(r, "id")
	err := cc.repo.Delete(carID)
	if err != nil {
		logger.Log().WithError(err).Error("Error during DeleteCar")
		switch v := err.(type) {
		case carerrors.CarError:
			// If car not found error.
			if v.Code == carerrors.NotFoundErrCode {
				render.Render(w, r, carerrors.HTTPErrNotFound(err))
				return
			}
		default:
			// If error due to other cases. This will be mostly internal server error
			render.Render(w, r, carerrors.HTTPErrInternalServer(err))
			return
		}
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
		render.Render(w, r, carerrors.HTTPErrInvalidRequest(err))
		return
	}

	id, err := cc.repo.Create(m.NewCar(cr))
	if err != nil {
		logger.Log().WithError(err).Error("Error during CreateCar")
		// If error due to other cases. This will be mostly internal server error
		render.Render(w, r, carerrors.HTTPErrInternalServer(err))
		return
	}
	response["id"] = id
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}

// GetAllCars is handler function to get all cars
func (cc CarsController) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := cc.repo.GetAll()
	if err != nil {
		logger.Log().WithError(err).Error("Error during GetAllCars")
		render.Render(w, r, carerrors.HTTPErrInternalServer(err))
		return
	}
	render.JSON(w, r, cars)
}

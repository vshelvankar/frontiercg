package main

import (
	"log"
	"net/http"

	r "github.com/vshelvankar/frontiercg/carmgmt/core/repository"

	c "github.com/vshelvankar/frontiercg/carmgmt/controller"
	d "github.com/vshelvankar/frontiercg/carmgmt/core/datastore"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// InitRoutes configuration
func InitRoutes(repo r.CarsRepository) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/api", func(router chi.Router) {
		router.Mount("/v1/cars", c.CarHandlerRoutes(repo))
	})

	return router
}

func main() {
	// init in memory datastore for cars
	r := d.NewCarsInMemoryDataStore()
	// Init routes and middlewares
	router := InitRoutes(r)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

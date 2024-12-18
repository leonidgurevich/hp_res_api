package routes

import (
	"github.com/gorilla/mux"
	"github.com/leonidgurevich/hc_rest_api/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/api/v1/users/{id}", controllers.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/v1/users", controllers.CreateUserHandler).Methods("POST")

	// Other routes
	// e.g., providers, payments, appointments, etc.
	router.HandleFunc("/api/v1/providers/{id}", controllers.GetProviderHandler).Methods("GET")
	// Add remaining routes...

	return router
}

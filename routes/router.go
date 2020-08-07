package router

import (
	"f1api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/circuits/", middleware.GetAllCircuits).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/drivers/", middleware.GetDrivers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/drivers/{id}", middleware.GetDriver).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/constructors/", middleware.GetConstructors).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/constructors/{id}", middleware.GetConstructor).Methods("GET", "OPTIONS")

	return router
}

package routes

import (
	"github.com/gorilla/mux"

	hdl "github.com/JuanVegaN/GoCalc.git/internal/v1/handlers"
)

func healtCheckRoutes(router *mux.Router) {

	router.HandleFunc("/healtcheck", hdl.HealtCheck).Methods("GET")

}

/*
* File: add.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This file contains the route localhost:3000/multiplication
*
 */

package routes

import (
	"github.com/gorilla/mux"

	hdl "github.com/JuanVegaN/GoCalc.git/internal/v1/handlers"
)

func multiplicationRoutes(router *mux.Router) {

	router.HandleFunc("/multiplication", hdl.MultiplicationHandler).Methods("POST")

}

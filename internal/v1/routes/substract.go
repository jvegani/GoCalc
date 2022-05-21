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
*   This file contains the route localhost:3000/substract
*
 */

package routes

import (
	hdl "github.com/JuanVegaN/GoCalc.git/internal/v1/handlers"
	"github.com/gorilla/mux"
)

func SubstractRoutes(router *mux.Router) {

	router.HandleFunc("/substract", hdl.SubstractHandler).Methods("POST")

}

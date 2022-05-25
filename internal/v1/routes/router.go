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
*   This file contains the route localhost:3000/index
*   And some types used in other files
*
 */

package routes

import (
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {

	// Add Routes
	addRoutes(router)

	// Division Routes
	divisionRoutes(router)

	// Mutiplication Routes
	multiplicationRoutes(router)

	// Sqrt Routes
	sqrtRoutes(router)

	//  Substract Routes
	SubstractRoutes(router)

	// HealtCheck
	healtCheckRoutes(router)

}

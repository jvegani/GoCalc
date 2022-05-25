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
*   This file contains the route localhost:3000/add
*
 */

package main

import (
	"log"
	"net/http"

	routes "github.com/JuanVegaN/GoCalc.git/internal/v1/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	routes.Router(router)

	log.Fatal(http.ListenAndServe(":3000", router))

}

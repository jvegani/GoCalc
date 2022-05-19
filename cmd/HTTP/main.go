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

	"github.com/JuanVegaN/GoCalc.git/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", routes.Index)
	router.HandleFunc("/add", routes.Add).Methods("GET")
	router.HandleFunc("/substract", routes.Substract).Methods("GET")
	router.HandleFunc("/multiplication", routes.Multiplication).Methods("GET")
	router.HandleFunc("/division", routes.Division).Methods("GET")
	router.HandleFunc("/sqrt", routes.Sqrt).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

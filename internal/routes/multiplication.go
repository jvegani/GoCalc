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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JuanVegaN/GoCalc.git/pkg/mathops"
)

func Multiplication(w http.ResponseWriter, r *http.Request) {
	var req JsonReq
	var resultAdd float64
	var ops string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error en  el JSON Request")
	}
	json.Unmarshal(reqBody, &req)
	if req.Float {
		ops = "Multiplicacion Flotantes"
		resultAdd = mathops.MultFloats(req.FirstOperator, req.SecondOperator)
	} else {
		resultAdd = float64(mathops.Mult(int(req.FirstOperator), int(req.SecondOperator)))
		ops = "Multiplication Enteros"
	}
	resp := JsonResp{
		Operation:      ops,
		FirstOperator:  req.FirstOperator,
		SecondOperator: req.SecondOperator,
		Result:         resultAdd,
		Message:        "La multiplicacion se ha realizado con exito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

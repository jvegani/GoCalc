package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	mathops "github.com/JuanVegaN/GoCalc.git/pkg/mathops/v1"
)

func SubstractHandler(w http.ResponseWriter, r *http.Request) {

	request := MathopsRequest{}
	var result float64

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err.Error())
		b, _ := json.Marshal(map[string]interface{}{"status": "failed", "error": "error with PayLoad Structure or Types"})
		http.Error(w, string(b), http.StatusUnprocessableEntity)
		return
	}

	if err := request.Validate(); err != nil {
		log.Println(err.Error())
		b, _ := json.Marshal(map[string]string{"status": "failed", "error": err.Error()})
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	if *request.NumOne == float64(int64(*request.NumOne)) && *request.NumTwo == float64(int64(*request.NumTwo)) { //the NumOne/Two are int
		result = float64(mathops.Substract(int(*request.NumOne), int(*request.NumTwo)))
	} else {
		result = mathops.SubstractFloats(*request.NumOne, *request.NumTwo)
	}

	if err := json.NewEncoder(w).Encode(
		map[string]interface{}{"status": "success", "substract_result": result},
	); err != nil {
		b, _ := json.Marshal(map[string]string{"status": "failed", "error": "error while writing http response"})
		http.Error(w, string(b), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

package handlers

import (
	"encoding/json"
	"net/http"

	mathops "github.com/JuanVegaN/GoCalc.git/pkg/mathops/v1"
	validator "github.com/go-playground/validator/v10"
)

func MultiplicationHandler(w http.ResponseWriter, r *http.Request) {
	request := MathopsRequest{}
	// Here Decode make the validation if that fields -Json to Struct- are numbers
	// the other way is use an interface{} for decode JSON in it and to can validate its fields.
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here validator/v10 make the validation if that fields are present in the Json
	val := validator.New()
	if err := val.Struct(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate result for int or float values
	var result float64

	if request.NumOne == float64(int64(request.NumOne)) && request.NumTwo == float64(int64(request.NumTwo)) { //the NumOne/Two are int
		result = float64(mathops.Mult(int(request.NumOne), int(request.NumTwo)))
	} else {
		result = mathops.MultFloats(request.NumOne, request.NumTwo)
	}

	// Write on HTTP response writer
	if err := json.NewEncoder(w).Encode(
		map[string]interface{}{"status": "success", "add_result": result},
	); err != nil {
		http.Error(w, "panic error while writing http response", http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

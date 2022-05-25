package handlers

import (
	"encoding/json"
	"net/http"

	mathops "github.com/JuanVegaN/GoCalc.git/pkg/mathops/v1"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: implement request validation (num, required)

	request := &AddRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := mathops.Add(request.NumOne, request.NumTwo)

	json.NewEncoder(w).Encode(
		map[string]int{"add_result": result},
	)

}

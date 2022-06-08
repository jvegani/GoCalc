package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	mathops "github.com/JuanVegaN/GoCalc.git/pkg/mathops/v1"
)

func SqrtHandler(w http.ResponseWriter, r *http.Request) {
	request := MathopsRequest{}

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

	if err := json.NewEncoder(w).Encode(
		map[string]interface{}{"status": "success", "sqrt_result": mathops.Sqrt(*request.NumOne)},
	); err != nil {
		b, _ := json.Marshal(map[string]string{"status": "failed", "error": "error while writing http response"})
		http.Error(w, string(b), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

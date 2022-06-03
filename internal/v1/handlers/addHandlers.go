package handlers

import (
	"encoding/json"
	"net/http"

	mathops "github.com/JuanVegaN/GoCalc.git/pkg/mathops/v1"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {

	/*
		TODO: implement request validation (num, required)

			R/ There are two ways I found: use Go Playground Validator package
			   or use an interface{} for decode JSON in it and to can validate its fields.
			   I use in this addHandler function the interface{} way but in the others
			   handlers I use the Go Playground Validator package because it is most scalable and reusable
	*/

	var request map[string]interface{} // instead of request := &AddRequest{}

	// Validation Request's Json Structure
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil { //Decode can be receive a Reference type (map,pointer or slice)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validation Required fields
	var num_one_ok, num_two_ok, num_int_number_ok, num_float_number_ok, num_not_number bool // false for default (zeroed value)
	for k, v := range request {
		if k == "num_one" {
			num_one_ok = true
		} else if k == "num_two" {
			num_two_ok = true
		} else {
			continue
		}
		if num_one_ok || num_two_ok {
			switch v.(type) {
			case int:
				num_int_number_ok = true
			case float64:
				num_float_number_ok = true
			default:
				num_not_number = true
			}
		}
	}
	if !(num_one_ok && num_two_ok) {
		http.Error(w, "num_one and num_two are required", http.StatusBadRequest)
		return
	} else if num_not_number {
		http.Error(w, "num_one and num_two must be numbers", http.StatusBadRequest)
		return
	}

	// Calculate result for int or float values
	var result interface{}
	if num_int_number_ok {
		result = mathops.Add(request["num_one"].(int), request["num_two"].(int)) // HERE if I put num_one int and num_two float IT'S work, why?
	} else if num_float_number_ok {
		result = mathops.AddFloats(request["num_one"].(float64), request["num_two"].(float64))
	}

	// Write on HTTP response writer
	if err := json.NewEncoder(w).Encode(
		map[string]interface{}{"status": "success", "add_result": result},
	); err != nil {
		http.Error(w, "panic error while writing http response", http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)

}

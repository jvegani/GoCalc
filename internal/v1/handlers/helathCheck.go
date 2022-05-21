package handlers

import (
	"encoding/json"
	"net/http"
)

func HealtCheck(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}

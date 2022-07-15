package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if v != nil {
		json.NewEncoder(w).Encode(v)
	} else {
		json.NewEncoder(w)
	}
}

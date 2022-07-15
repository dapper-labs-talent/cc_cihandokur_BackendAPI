package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ApiError struct{}

func (e ApiError) HandleErr(w http.ResponseWriter, status int, message string) {

	log.Println(message)

	err := make(map[string]string)

	err["Message"] = message
	err["Status"] = strconv.Itoa(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

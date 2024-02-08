package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	RespondStatus(w, code)
	w.Write(data)
}

func RespondStatus(w http.ResponseWriter, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
}

func RespondErrorMessage(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5XX error:", msg)
	}
	RespondJSON(w, code, struct {
		Error string `json:"error"`
	}{Error: msg})
}

func CheckAndRespondError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Fatal("There was a server error:", err)
		RespondErrorMessage(w, http.StatusInternalServerError, fmt.Sprintf("Server error: %s", err))
	}
}

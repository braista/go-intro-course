package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"
)

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
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

func RespondError(w http.ResponseWriter, err error) {
	log.Println("there was a server error:", err)
	RespondErrorMessage(w, http.StatusInternalServerError, fmt.Sprintf("there was a server error: %s", err))
}

func RespondDBError(w http.ResponseWriter, err error) {
	if dbError, ok := err.(*pq.Error); ok {
		log.Printf("there was a db error (%s): %s", dbError.Code, dbError)
		switch dbError.Code.Name() {
		case "foreign_key_violation", "unique_violation":
			RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		default:
			RespondError(w, err)
		}
		return
	}
	if err == sql.ErrNoRows {
		RespondStatus(w, http.StatusNotFound)
		return
	}
}

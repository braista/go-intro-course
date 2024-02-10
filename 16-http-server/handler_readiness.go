package main

import (
	"net/http"
)

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "OK"})
}

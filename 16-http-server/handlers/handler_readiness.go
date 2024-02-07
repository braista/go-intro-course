package handlers

import (
	"net/http"

	"github.com/braista/go-intro-course/16-http-server/utils"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "OK"})
}

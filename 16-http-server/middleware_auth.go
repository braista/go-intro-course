package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/braista/go-intro-course/16-http-server/internal/auth"
	"github.com/braista/go-intro-course/16-http-server/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			RespondErrorMessage(w, http.StatusForbidden, fmt.Sprintf("auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetCurrentUser(r.Context(), apiKey)
		if err != nil && err == sql.ErrNoRows {
			RespondStatus(w, http.StatusForbidden)
			return
		}
		handler(w, r, user)
	}
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/braista/go-intro-course/16-http-server/internal/auth"
	"github.com/braista/go-intro-course/16-http-server/internal/database"
	"github.com/braista/go-intro-course/16-http-server/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondErrorMessage(w, http.StatusForbidden, fmt.Sprintf("auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetCurrentUser(r.Context(), apiKey)
		if err != nil {
			utils.RespondError(w, err)
			return
		}
		handler(w, r, user)
	}
}

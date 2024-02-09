package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/braista/go-intro-course/16-http-server/internal/database"
	"github.com/braista/go-intro-course/16-http-server/models"
	"github.com/braista/go-intro-course/16-http-server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const UsersPath string = "/users"

func (apiCfg *ApiConfig) HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetUsers(r.Context())
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, models.DatabaseUsersToUsers(users))
}

func (apiCfg *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerAddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := userParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	utils.RespondJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Printf("Updating user with id %s", id)
	decoder := json.NewDecoder(r.Body)
	params := userParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	rows, err := apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:        uuid.MustParse(id),
		Name:      params.Name,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	if rows == 0 {
		utils.RespondErrorMessage(w, http.StatusNotFound, fmt.Sprintf("User with id %s doesn't exist", id))
		return
	}
	utils.RespondStatus(w, http.StatusNoContent)
}

func (apiCfg *ApiConfig) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	// getting pathparam with chi
	id := chi.URLParam(r, "id")
	log.Printf("Deleting user with id %s", id)
	rows, err := apiCfg.DB.DeleteUser(r.Context(), uuid.MustParse(id))
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	if rows == 0 {
		utils.RespondErrorMessage(w, http.StatusNotFound, fmt.Sprintf("User with id %s doesn't exist", id))
		return
	}
	utils.RespondStatus(w, http.StatusOK)
}

type userParameters struct {
	Name string `json:"name"`
}

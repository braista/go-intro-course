package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/braista/go-intro-course/16-http-server/internal/database"
	"github.com/braista/go-intro-course/16-http-server/models"
	"github.com/braista/go-intro-course/16-http-server/utils"
	"github.com/google/uuid"
)

const FeedsPath string = "/feeds"

type feedParameters struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (apiCfg *ApiConfig) HandleAddFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := feedParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondDBError(w, err)
		return
	}
	utils.RespondJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/braista/go-intro-course/16-http-server/internal/database"
	"github.com/braista/go-intro-course/16-http-server/models"
	"github.com/braista/go-intro-course/16-http-server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (apiCfg *ApiConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondDBError(w, err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}

func (apiCfg *ApiConfig) HandleGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		utils.RespondDBError(w, err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}

func (apiCfg *ApiConfig) HandleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID := chi.URLParam(r, "id")
	log.Printf("[%s] generating follow to feed '%s'", user.ID, feedID)
	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		utils.RespondErrorMessage(w, http.StatusBadRequest, "feed uuid is invalid")
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feedUUID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
	})
	if dbError, ok := err.(*pq.Error); ok {
		switch dbError.Constraint {
		case "feeds_follows_feed_id_fkey":
			utils.RespondErrorMessage(w, http.StatusBadRequest, "feed doesn't exist")
		case "feeds_follows_user_id_feed_id_key":
			utils.RespondErrorMessage(w, http.StatusBadRequest, "feed already followed")
		default:
			utils.RespondError(w, err)
		}
		return
	}
	if err != nil {
		utils.RespondError(w, err)
		return
	}
	log.Printf("[%s] feed '%s' successfully followed", user.ID, feedID)
	utils.RespondJSON(w, http.StatusCreated, models.DatabaseFeedFollowtoFeedFollow(feedFollow))
}

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/braista/go-intro-course/16-http-server/internal/database"
	"github.com/braista/go-intro-course/16-http-server/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

const feedsPath string = "/feeds"

type FeedHandler interface {
	handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) models.User
	handleGetFeeds(w http.ResponseWriter, r *http.Request) []models.User
	handleGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) []models.Feed
	handleFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) models.FeedFollow
	HandleUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) models.FeedFollow
}

type feedParameters struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := feedParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondError(w, err)
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
		RespondDBError(w, err)
		return
	}
	RespondJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondDBError(w, err)
		return
	}
	RespondJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}

func (apiCfg *apiConfig) handleGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		RespondDBError(w, err)
		return
	}
	RespondJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}

func (apiCfg *apiConfig) handleFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID := chi.URLParam(r, "id")
	log.Printf("[%s] generating follow to feed '%s'", user.ID, feedID)
	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		RespondErrorMessage(w, http.StatusBadRequest, "feed uuid is invalid")
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
			RespondErrorMessage(w, http.StatusBadRequest, "feed doesn't exist")
		case "feeds_follows_user_id_feed_id_key":
			RespondErrorMessage(w, http.StatusBadRequest, "feed already followed")
		default:
			RespondError(w, err)
		}
		return
	}
	if err != nil {
		RespondError(w, err)
		return
	}
	log.Printf("[%s] feed '%s' successfully followed", user.ID, feedID)
	RespondJSON(w, http.StatusCreated, models.DatabaseFeedFollowtoFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) HandleUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID := chi.URLParam(r, "id")
	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		RespondErrorMessage(w, http.StatusBadRequest, "feed uuid is invalid")
		return
	}
	feed, err := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feedUUID,
		UserID: user.ID,
	})
	if dbError, ok := err.(*pq.Error); ok {
		log.Println("wut: " + dbError.Constraint)
		switch dbError.Constraint {
		case "feeds_follows_feed_id_fkey":
			RespondErrorMessage(w, http.StatusBadRequest, "feed doesn't exist")
		default:
			RespondError(w, err)
		}
		return
	}
	if err == sql.ErrNoRows {
		RespondErrorMessage(w, http.StatusBadRequest, "feed not followed")
		return
	}
	if err != nil {
		RespondError(w, err)
		return
	}
	log.Printf("[%s] feed '%s' successfully unfollowed", user.ID, feedID)
	RespondJSON(w, http.StatusOK, feed)
}

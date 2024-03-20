package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Popo422/rssagg/internal/auth"
	"github.com/Popo422/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respWithError(w, 400, "Something Went wrong parsing the json")
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Failed to create a user %v", err))
	}

	respWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Something Went Wrong querying feeds %v", err))
		return
	}

	respWithJSON(w, 201, databaseFeedsToFeeds(feeds))

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Popo422/rssagg/internal/auth"
	"github.com/Popo422/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respWithError(w, 400, "Something Went wrong parsing the json")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Failed to create a user %v", err))
	}

	respWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respWithJSON(w, 200, databaseUserToUser(user))
}

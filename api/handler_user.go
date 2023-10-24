package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Jhnvlglmlbrt/rss-aggregator/internal/database"
	"github.com/Jhnvlglmlbrt/rss-aggregator/models"
	"github.com/Jhnvlglmlbrt/rss-aggregator/utils"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create a user: %v", err))
		return
	}

	utils.RespondWithJson(w, 201, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJson(w, 200, models.DatabaseUserToUser(user))
}

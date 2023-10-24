package api

import (
	"fmt"
	"net/http"

	"github.com/Jhnvlglmlbrt/rss-aggregator/internal/auth"
	"github.com/Jhnvlglmlbrt/rss-aggregator/internal/database"
	"github.com/Jhnvlglmlbrt/rss-aggregator/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(Handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Couldn't get apiKey: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPiKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		Handler(w, r, user)
	}
}

package api

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/rss-aggregator/usecase"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	usecase.RespondWithJson(w, 200, struct{}{})
}

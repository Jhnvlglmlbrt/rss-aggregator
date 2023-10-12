package api

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/rss-aggregator/usecase"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	usecase.RespondWithError(w, 400, "Something went wrong")
}

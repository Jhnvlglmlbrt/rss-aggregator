package api

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/rss-aggregator/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}

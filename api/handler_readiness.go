package api

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/rss-aggregator/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}

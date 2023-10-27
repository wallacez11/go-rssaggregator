package handlers

import (
	"net/http"

	json "github.com/wallacez11/go-rssaggregator/util"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	json.RespondWithJson(w, 200, struct{}{})
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	json.RespondWithError(w, 400, "Something went wrong")
}

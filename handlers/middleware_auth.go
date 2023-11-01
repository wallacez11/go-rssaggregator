package handlers

import (
	"fmt"
	"net/http"

	"github.com/wallacez11/go-rssaggregator/internal/auth"
	"github.com/wallacez11/go-rssaggregator/internal/database"
	utils "github.com/wallacez11/go-rssaggregator/util"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (ApiConfig *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 400, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := ApiConfig.Db.GetUserByApiKey(r.Context(), apikey)

		if err != nil {
			utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}

}

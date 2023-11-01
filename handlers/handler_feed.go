package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
	"github.com/wallacez11/go-rssaggregator/internal/database"
	utils "github.com/wallacez11/go-rssaggregator/util"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		Name string `json:"name" validate:"required"`
		URL  string `json:"url" validate:"required"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &Parameters{}
	err := decoder.Decode(params)
	if err != nil {
		fmt.Println("Error on decode JSON:", err)
	}

	validate := validator.New()
	errvalidation := validate.Struct(params)

	if errvalidation != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error on decode body: %v", errvalidation))
		return
	}

	feed, err := apiCfg.Db.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error creating a feed: %v ", err))
		return
	}

	utils.RespondWithJson(w, 201, utils.DatabaseConvertFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.Db.GetFeeds(r.Context())

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error creating a feed: %v ", err))
		return
	}

	utils.RespondWithJson(w, 201, utils.DatabaseMultipleFeeds(feeds))
}

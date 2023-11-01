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

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		FeedId uuid.UUID `json:"feed_id" validate:"required"`
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

	FeedFollow, err := apiCfg.Db.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error creating a feed: %v ", err))
		return
	}

	utils.RespondWithJson(w, 201, utils.DatabaseConvertFeedFollow(FeedFollow))
}

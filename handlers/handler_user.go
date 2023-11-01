package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wallacez11/go-rssaggregator/internal/database"
	utils "github.com/wallacez11/go-rssaggregator/util"
)

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &Parameters{}
	err := decoder.Decode(params)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error on decode body: %v", err))
		return
	}

	user, err := apiCfg.Db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error creating a user: %v ", err))
		return
	}

	utils.RespondWithJson(w, 201, utils.DatabaseConvertUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJson(w, 200, utils.DatabaseConvertUser(user))
}

package handlers

import (
	_ "github.com/lib/pq"
	"github.com/wallacez11/go-rssaggregator/internal/database"
)

type ApiConfig struct {
	Db *database.Queries
}

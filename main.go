package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	rss "github.com/wallacez11/go-rssaggregator/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	db "github.com/wallacez11/go-rssaggregator/handlers"
	"github.com/wallacez11/go-rssaggregator/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Could not load env variable")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("Could not load DB variable")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	apiCfg := db.ApiConfig{
		Db: database.New(conn),
	}

	v1Router := chi.NewRouter()
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"*"},
		MaxAge:           300,
	}))

	v1Router.Get("/ready", rss.HandlerReadiness)
	v1Router.Get("/err", rss.HandlerError)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeed)
	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	errs := srv.ListenAndServe()

	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Println("Port: ", portString)
}

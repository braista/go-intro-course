package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/braista/go-intro-course/16-http-server/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	// loading .env file to the environment
	godotenv.Load(".env")
	// database configuration
	dbURL := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	queries := database.New(connection)
	apiCfg := apiConfig{
		DB: queries,
	}
	// reading environment variables using os.Getenv()
	port := os.Getenv("PORT")
	log.Println("starting server on port", port)

	router := chi.NewRouter()
	// health-check handler
	router.Get("/healthz", handleHealthCheck)
	// users handler
	router.Get(usersPath, apiCfg.handlerGetUsers)
	router.Post(usersPath, apiCfg.handlerAddUser)
	router.Get(usersPath+"/current", apiCfg.MiddlewareAuth(apiCfg.handleGetUser))
	router.Patch(usersPath+"/users/{id}", apiCfg.handleUpdateUser)
	router.Delete(usersPath+"/users/{id}", apiCfg.handlerDeleteUser)
	//feeds handler
	router.Post(feedsPath, apiCfg.MiddlewareAuth(apiCfg.handleAddFeed))
	router.Get(feedsPath, apiCfg.handleGetFeeds)
	router.Get(usersPath+"/current"+feedsPath, apiCfg.MiddlewareAuth(apiCfg.handleGetUserFeeds))
	router.Get(feedsPath+"/{id}/follow", apiCfg.MiddlewareAuth(apiCfg.handleCreateFeedFollow))
	router.Get(feedsPath+"/{id}/unfollow", apiCfg.MiddlewareAuth(apiCfg.HandleUnfollowFeed))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type apiConfig struct {
	DB *database.Queries
}

package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/braista/go-intro-course/16-http-server/handlers"
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
	apiCfg := handlers.ApiConfig{
		DB: queries,
	}
	// reading environment variables using os.Getenv()
	port := os.Getenv("PORT")
	log.Println("starting server on port", port)

	router := chi.NewRouter()
	// health-check handler
	router.Get("/healthz", handlers.HandleHealthCheck)
	// users handler
	router.Get(handlers.UsersPath, apiCfg.HandlerGetUsers)
	router.Post(handlers.UsersPath, apiCfg.HandlerAddUser)
	router.Get(handlers.UsersPath+"/current", apiCfg.MiddlewareAuth(apiCfg.HandleGetUser))
	router.Patch(handlers.UsersPath+"/users/{id}", apiCfg.HandleUpdateUser)
	router.Delete(handlers.UsersPath+"/users/{id}", apiCfg.HandlerDeleteUser)
	//feeds handler
	router.Post(handlers.FeedsPath, apiCfg.MiddlewareAuth(apiCfg.HandleAddFeed))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

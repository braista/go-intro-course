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
	router.Get("/healthz", handlers.HandleHealthCheck)
	router.Get("/users", apiCfg.HandlerGetUsers)
	router.Post("/users", apiCfg.HandlerAddUser)
	router.Patch("/users/{id}", apiCfg.HandleUpdateUser)
	router.Delete("/users/{id}", apiCfg.HandlerDeleteUser)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/braista/go-intro-course/16-http-server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// loading .env file to the environment
	godotenv.Load(".env")
	// reading environment variables using os.Getenv()
	port := os.Getenv("PORT")
	log.Println("starting server on port", port)
	router := chi.NewRouter()
	router.Get("/healthz", handlers.HandleHealthCheck)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

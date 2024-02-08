package handlers

import "github.com/braista/go-intro-course/16-http-server/internal/database"

type ApiConfig struct {
	DB *database.Queries
}

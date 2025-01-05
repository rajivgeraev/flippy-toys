package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/config"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/middleware"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/handler"
)

func main() {
	cfg := config.LoadConfig()
	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)

	userHandler := handler.NewUserHandler(cfg)

	r.HandleFunc("/api/v1/users/me", userHandler.GetMe).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/auth/validate", userHandler.ValidateUser).Methods("POST", "OPTIONS")

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

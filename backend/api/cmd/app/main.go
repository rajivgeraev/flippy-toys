// cmd/app/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/config"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/database"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/middleware"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/handler"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/repository/postgres"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/service"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()
	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is required")
	}
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	// Инициализация БД
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Инициализация зависимостей
	userRepo := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(cfg, userService)

	// Настройка роутера
	r := mux.NewRouter()

	// Middleware для всех запросов
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)

	// Открытые маршруты
	r.HandleFunc("/api/v1/auth/validate", userHandler.ValidateUser).
		Methods("POST", "OPTIONS")

	// Защищенные маршруты
	protected := r.PathPrefix("/api/v1").Subrouter()
	protected.Use(middleware.TelegramAuth(cfg.BotToken))

	protected.HandleFunc("/users/me", userHandler.GetMe).
		Methods("GET", "OPTIONS")
	protected.HandleFunc("/users/phone", userHandler.UpdatePhone).
		Methods("POST", "OPTIONS")

	// Запуск сервера
	port := ":" + cfg.Port
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

// r.HandleFunc("/api/v1/users/me", userHandler.GetMe).Methods("GET", "OPTIONS")
// r.HandleFunc("/api/v1/auth/validate", userHandler.ValidateUser).Methods("POST", "OPTIONS")

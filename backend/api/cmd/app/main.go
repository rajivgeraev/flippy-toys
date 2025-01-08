package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/cloudinary"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/config"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/database"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/middleware"
	toyHandler "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/handler"
	toyRepo "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/repository/postgres"
	toyService "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/service"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/handler"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/repository/postgres"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/service"
)

var emptyHandleFunc = func(w http.ResponseWriter, r *http.Request) {}

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()
	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is required")
	}
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	if cfg.CloudinaryName == "" || cfg.CloudinaryApiKey == "" || cfg.CloudinaryApiSecret == "" || cfg.CloudinaryUploadPreset == "" {
		log.Fatal("Cloudinary configuration is required")
	}

	// Инициализация БД
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	cloudinaryConfig := cloudinary.NewConfig(
		cfg.CloudinaryName,
		cfg.CloudinaryApiKey,
		cfg.CloudinaryApiSecret,
		cfg.CloudinaryUploadPreset,
	)

	// Инициализация Cloudinary
	cloudinaryClient, err := cloudinary.NewClient(cloudinaryConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	// Инициализация репозиториев
	userRepo := postgres.NewUserRepository(db)
	toyRepo := toyRepo.NewToyRepository(db)

	// Инициализация сервисов
	userService := service.NewUserService(userRepo)
	toyService := toyService.NewToyService(toyRepo, cloudinaryClient)

	// Инициализация хендлеров
	userHandler := handler.NewUserHandler(cfg, userService)
	toysHandler := toyHandler.NewToyHandler(toyService)

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
	protected.Use(middleware.TelegramAuth(cfg.BotToken, userService))

	// User routes
	protected.HandleFunc("/users/me", userHandler.GetMe).
		Methods("GET", "OPTIONS")
	protected.HandleFunc("/users/phone", userHandler.UpdatePhone).
		Methods("POST", "OPTIONS")

	// Toy routes
	protected.HandleFunc("/toys", toysHandler.CreateToy).
		Methods("POST", "OPTIONS")
	protected.HandleFunc("/toys/id/{id}", emptyHandleFunc).
		Methods("OPTIONS")
	protected.HandleFunc("/toys/id/{id}", toysHandler.GetToy).
		Methods("GET")
	protected.HandleFunc("/toys/id/{id}", toysHandler.UpdateToy).
		Methods("PUT")

	protected.HandleFunc("/toys/my", toysHandler.GetUserToys).
		Methods("GET", "OPTIONS")

	protected.HandleFunc("/toys/upload/params", toysHandler.GetUploadParams).
		Methods("GET", "OPTIONS")

	// Запуск сервера
	port := ":" + cfg.Port
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

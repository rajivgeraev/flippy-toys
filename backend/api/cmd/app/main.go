package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	childHandlerPkg "github.com/rajivgeraev/flippy-toys/backend/api/internal/child/handler"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/cloudinary"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/config"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/database"
	customMiddleware "github.com/rajivgeraev/flippy-toys/backend/api/internal/common/middleware"
	toyHandler "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/handler"
	toyRepo "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/repository/postgres"
	toyService "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/service"
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
	r := chi.NewRouter()

	// Middleware для всех запросов
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.CORS)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// API v1 router
	r.Route("/api/v1", func(r chi.Router) {
		// Auth routes
		r.Route("/auth", func(r chi.Router) {
			r.Post("/validate", userHandler.ValidateUser)
		})

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.TelegramAuth(cfg.BotToken, userService))

			// User routes
			r.Route("/users", func(r chi.Router) {
				r.Get("/me", userHandler.GetMe)
				r.Post("/phone", userHandler.UpdatePhone)
				r.Get("/{id}", userHandler.GetUser)
			})

			// Toy routes
			r.Route("/toys", func(r chi.Router) {
				r.Get("/my", toysHandler.GetUserToys)
				r.Get("/upload/params", toysHandler.GetUploadParams)
				r.Get("/{id}", toysHandler.GetToy)
				r.Post("/{id}", toysHandler.UpdateToy)
				r.Get("/", toysHandler.ListToys)
				r.Post("/", toysHandler.CreateToy)
			})

			// Children routes
			r.Route("/children", func(r chi.Router) {
				childHandler := childHandlerPkg.NewChildHandler()
				r.Get("/", childHandler.GetChildren)
			})
		})
	})

	// Запуск сервера
	port := ":" + cfg.Port
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

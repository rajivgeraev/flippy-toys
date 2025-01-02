package main

import (
    "log"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/config"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/handler"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/repository/postgres"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/service"
)

func main() {
    // Загружаем конфигурацию
    cfg, err := config.New()
    if err != nil {
        log.Fatalf("Config error: %s", err)
    }

    // Инициализируем подключение к БД
    db, err := postgres.NewDB(&cfg.PostgresConfig)
    if err != nil {
        log.Fatalf("DB error: %s", err)
    }
    defer db.Close()

    // Инициализируем репозитории
    userRepo := postgres.NewUserRepository(db)

    // Инициализируем сервисы
    deps := service.Deps{
        UserRepo: userRepo,
    }
    services := service.NewServices(deps)

    // Инициализируем хендлеры
    handlers := handler.NewHandler(services)

    // Запускаем сервер
    router := handlers.InitRoutes()
    
    log.Printf("Starting server on port %s", cfg.ServerConfig.Port)
    if err := router.Run(":" + cfg.ServerConfig.Port); err != nil {
        log.Fatalf("Server error: %s", err)
    }
}
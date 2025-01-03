// cmd/app/main.go

package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/handler"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/middleware"
)

func main() {
    r := mux.NewRouter()
    
    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.CORS)
    
    userHandler := handler.NewUserHandler()
    r.HandleFunc("/api/v1/users/me", userHandler.GetMe).Methods("GET", "OPTIONS")
    
    port := ":8080"
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(port, r))
}
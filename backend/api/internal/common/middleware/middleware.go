// internal/common/middleware/middleware.go
package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/rajivgeraev/flippy-toys/backend/api/internal/auth/telegram"
)

// CORS middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Telegram-Init-Data")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

// TelegramAuth middleware
func TelegramAuth(botToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			initData := r.Header.Get("X-Telegram-Init-Data")
			if initData == "" {
				http.Error(w, "No auth data provided", http.StatusUnauthorized)
				return
			}

			telegramData, err := telegram.ValidateInitData(initData, botToken)
			if err != nil {
				http.Error(w, "Invalid auth data", http.StatusUnauthorized)
				return
			}

			// Добавляем telegram_id в контекст запроса
			ctx := context.WithValue(r.Context(), "telegram_id", telegramData.User.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

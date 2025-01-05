// internal/common/middleware/middleware.go
package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rajivgeraev/flippy-toys/backend/api/internal/auth/telegram"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/model"
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
func TelegramAuth(botToken string, userService UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("=== TelegramAuth Middleware ===")
			fmt.Printf("Request URL: %s", r.URL.Path)
			fmt.Printf("Request Method: %s", r.Method)

			initData := r.Header.Get("X-Telegram-Init-Data")
			if initData == "" {
				fmt.Printf("No auth data provided")
				http.Error(w, "No auth data provided", http.StatusUnauthorized)
				return
			}

			telegramData, err := telegram.ValidateInitData(initData, botToken)
			if err != nil {
				fmt.Printf("Invalid auth data: %v", err)
				http.Error(w, "Invalid auth data", http.StatusUnauthorized)
				return
			}

			fmt.Printf("Telegram User ID: %d", telegramData.User.ID)

			// Получаем пользователя из БД
			user, err := userService.GetUserByTelegramID(telegramData.User.ID)
			if err != nil {
				fmt.Printf("Error getting user: %v", err)
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			fmt.Printf("User UUID: %s", user.ID)

			ctx := context.WithValue(r.Context(), "telegram_id", telegramData.User.ID)
			ctx = context.WithValue(ctx, "user_id", user.ID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type UserService interface {
	GetUserByTelegramID(telegramID int64) (*model.User, error)
}

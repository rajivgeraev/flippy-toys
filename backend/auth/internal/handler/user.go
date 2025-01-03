// internal/handler/user.go

package handler

import (
   "encoding/json"
   "net/http"
   "time"
   "log"

   "github.com/rajivgeraev/flippy-toys/backend/auth/internal/model"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
   return &UserHandler{}
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
    mockUser := model.User{
        TelegramID: 123456789,
        Username:   "johndoe",
        FirstName:  "John",
        LastName:   "Doe",
        IsPremium:  true,
        PhotoURL:   "https://t.me/i/userpic/320/example.jpg",
        CreatedAt:  time.Now().Add(-24 * time.Hour),
        LastLogin:  time.Now(),
    }

    w.Header().Set("Content-Type", "application/json")
    
    if err := json.NewEncoder(w).Encode(mockUser); err != nil {
        log.Printf("Error encoding response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    log.Printf("Successfully returned user data for TelegramID: %d", mockUser.TelegramID)
}

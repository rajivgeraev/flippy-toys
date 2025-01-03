package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rajivgeraev/flippy-toys/backend/auth/internal/config"
	"github.com/rajivgeraev/flippy-toys/backend/auth/internal/model"
	"github.com/rajivgeraev/flippy-toys/backend/auth/internal/telegram"
)

type UserHandler struct {
	cfg *config.Config
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	return &UserHandler{
		cfg: cfg,
	}
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

// В ValidateUser используем конфиг:
func (h *UserHandler) ValidateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InitData string `json:"init_data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userData, err := telegram.ValidateInitData(req.InitData, h.cfg.BotToken)
	if err != nil {
		log.Printf("Error validating initData: %v", err)
		http.Error(w, "Invalid Telegram data", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

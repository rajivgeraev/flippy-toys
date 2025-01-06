// internal/user/handler/user.go
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/config"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/service"
)

type UserHandler struct {
	cfg     *config.Config
	service *service.UserService
}

func NewUserHandler(cfg *config.Config, service *service.UserService) *UserHandler {
	return &UserHandler{
		cfg:     cfg,
		service: service,
	}
}

// ValidateUser обрабатывает первичную аутентификацию через Telegram
func (h *UserHandler) ValidateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InitData string `json:"init_data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Telegram InitData : %s\n", req.InitData)

	user, err := h.service.ProcessTelegramAuth(req.InitData, h.cfg.BotToken)
	if err != nil {
		log.Printf("Error processing telegram auth: %v", err)
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdatePhone обновляет номер телефона пользователя
func (h *UserHandler) UpdatePhone(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TelegramID int64  `json:"telegram_id"`
		Phone      string `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.UpdatePhone(req.TelegramID, req.Phone)
	if err != nil {
		log.Printf("Error updating phone: %v", err)
		http.Error(w, "Failed to update phone", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetMe возвращает информацию о текущем пользователе
func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	// Получаем telegram_id из контекста (установленного middleware)
	telegramID := r.Context().Value("telegram_id").(int64)

	fmt.Printf("GetToysByUserID Start : %s\n", "telegramID")

	user, err := h.service.GetUserByTelegramID(telegramID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Вспомогательные структуры для ответов
type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *UserHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func (h *UserHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

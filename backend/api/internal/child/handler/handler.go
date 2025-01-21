// backend/api/internal/child/handler/handler.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/child/models"
)

type ChildHandler struct{}

func NewChildHandler() *ChildHandler {
	return &ChildHandler{}
}

func (h *ChildHandler) GetChildren(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(uuid.UUID)
	if !ok {
		fmt.Printf("No user_id in context\n")
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Mock data for test user
	if userID.String() == "7ee06eb2-71d4-4c5d-87c1-5619f0f79aba" {
		children := []models.Child{
			{
				ID:        "child-123",
				ParentID:  userID.String(),
				Name:      "Миша",
				Age:       7,
				Gender:    "male",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        "child-124",
				ParentID:  userID.String(),
				Name:      "Виктория",
				Age:       5,
				Gender:    "female",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(children)
		return
	}

	// Для остальных пользователей возвращаем пустой массив
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]models.Child{})
}

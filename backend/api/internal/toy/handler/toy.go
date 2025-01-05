// internal/toy/handler/toy.go
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/service"
)

type ToyHandler struct {
	service *service.ToyService
}

func NewToyHandler(service *service.ToyService) *ToyHandler {
	return &ToyHandler{service: service}
}

func (h *ToyHandler) CreateToy(w http.ResponseWriter, r *http.Request) {
	// Получаем multipart form с фотографиями
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		AgeMin      int    `json:"age_min"`
		AgeMax      int    `json:"age_max"`
		Condition   string `json:"condition"`
		Category    string `json:"category"`
	}

	if err := json.Unmarshal([]byte(r.MultipartForm.Value["data"][0]), &input); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Собираем фотографии
	var photos [][]byte
	files := r.MultipartForm.File["photos"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to read photo", http.StatusBadRequest)
			return
		}
		defer file.Close()

		photoData := make([]byte, fileHeader.Size)
		if _, err := file.Read(photoData); err != nil {
			http.Error(w, "Failed to read photo", http.StatusBadRequest)
			return
		}

		photos = append(photos, photoData)
	}

	// Получаем пользователя из контекста
	userID := r.Context().Value("user_id").(uuid.UUID)

	toy, err := h.service.CreateToy(r.Context(), service.CreateToyInput{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		AgeRange: model.AgeRange{
			Min: input.AgeMin,
			Max: input.AgeMax,
		},
		Condition: model.ToyCondition(input.Condition),
		Category:  model.ToyCategory(input.Category),
		Photos:    photos,
	})

	if err != nil {
		http.Error(w, "Failed to create toy", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toy)
}

func (h *ToyHandler) GetToy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toyID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid toy ID", http.StatusBadRequest)
		return
	}

	toy, err := h.service.GetToy(r.Context(), toyID)
	if err != nil {
		http.Error(w, "Failed to get toy", http.StatusInternalServerError)
		return
	}

	if toy == nil {
		http.Error(w, "Toy not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toy)
}

func (h *ToyHandler) GetUserToys(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("=== GetUserToys Handler ===\n")

	userID, ok := r.Context().Value("user_id").(uuid.UUID)
	if !ok {
		fmt.Printf("No user_id in context\n")
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	fmt.Printf("Getting toys for user: %s\n", userID)

	toys, err := h.service.GetToysByUserID(userID)
	if err != nil {
		fmt.Printf("Error getting toys: %v\n", err)
		http.Error(w, "Failed to get toys", http.StatusInternalServerError)
		return
	}

	if toys == nil {
		fmt.Printf("No toys found, returning empty array\n")
		toys = []model.Toy{}
	} else {
		fmt.Printf("Found %d toys\n", len(toys))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(toys); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

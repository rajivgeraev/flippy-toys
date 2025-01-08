// internal/toy/handler/toy.go
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	toyModel "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
	toySrv "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/service"
)

type ToyHandler struct {
	service *toySrv.ToyService
}

func NewToyHandler(service *toySrv.ToyService) *ToyHandler {
	return &ToyHandler{service: service}
}

func (h *ToyHandler) CreateToy(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string                   `json:"title"`
		Description string                   `json:"description"`
		Condition   string                   `json:"condition"`
		Category    string                   `json:"category"`
		Photos      []toySrv.CloudinaryPhoto `json:"photos"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(uuid.UUID)

	condition := toyModel.ToyCondition(input.Condition)
	if condition == "" {
		http.Error(w, "Invalid request body - condition is required ", http.StatusBadRequest)
	}
	category := toyModel.ToyCategory(input.Category)
	if category == "" {
		http.Error(w, "Invalid request body - category is required ", http.StatusBadRequest)
	}

	toy, err := h.service.CreateToy(r.Context(), toySrv.CreateToyInput{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   &condition,
		Category:    &category,
		Photos:      input.Photos,
	})

	if err != nil {
		log.Printf("Error creating toy: %v", err)
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
	fmt.Printf(">>==GetUserToys Start : %s\n", "userID")

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
		toys = []toyModel.Toy{}
	} else {
		fmt.Printf("Found %d toys\n", len(toys))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(toys); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *ToyHandler) ListActive(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры пагинации
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	toys, err := h.service.ListToys(r.Context(), page, pageSize)
	if err != nil {
		log.Printf("Error listing toys: %v", err)
		http.Error(w, "Failed to list toys", http.StatusInternalServerError)
		return
	}

	if toys == nil {
		toys = []toyModel.Toy{} // Возвращаем пустой массив вместо nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toys)
}

func (h *ToyHandler) GetUploadParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params, err := h.service.GetUploadParams(ctx)
	if err != nil {
		log.Printf("Error getting upload params: %v", err)
		http.Error(w, "Failed to get upload parameters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(params)
}

func (h *ToyHandler) UpdateToy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toyID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid toy ID", http.StatusBadRequest)
		return
	}

	var input toySrv.UpdateToyInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	toy, err := h.service.UpdateToy(r.Context(), toyID, input)
	if err != nil {
		http.Error(w, "Failed to update toy", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toy)
}

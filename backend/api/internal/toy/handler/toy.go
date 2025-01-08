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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	log.Printf("Received form data: %+v", r.MultipartForm.Value)

	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		AgeMin      *int    `json:"age_min,omitempty"`
		AgeMax      *int    `json:"age_max,omitempty"`
		Condition   *string `json:"condition,omitempty"`
		Category    *string `json:"category,omitempty"`
	}

	if err := json.Unmarshal([]byte(r.MultipartForm.Value["data"][0]), &input); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Валидация обязательных полей
	if input.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
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

	userID := r.Context().Value("user_id").(uuid.UUID)

	// Создаем input для сервиса
	createInput := service.CreateToyInput{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Photos:      photos,
	}

	fmt.Printf("\n>>== createInput : %+v\n", createInput)

	// Добавляем опциональные поля если они есть
	if input.AgeMin != nil && input.AgeMax != nil {
		createInput.AgeRange = &model.AgeRange{
			Min: *input.AgeMin,
			Max: *input.AgeMax,
		}
	}

	if input.Condition != nil {
		condition := model.ToyCondition(*input.Condition)
		createInput.Condition = &condition
	}

	if input.Category != nil {
		category := model.ToyCategory(*input.Category)
		createInput.Category = &category
	}
	fmt.Printf(">>== Start  1: %s\n", "userID")

	toy, err := h.service.CreateToy(r.Context(), createInput)
	if err != nil {
		fmt.Printf("\n>>== Error : %w\n", err)

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
		toys = []model.Toy{}
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
		toys = []model.Toy{} // Возвращаем пустой массив вместо nil
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

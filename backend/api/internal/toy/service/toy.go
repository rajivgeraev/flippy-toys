// internal/toy/service/toy.go
package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/cloudinary"
	toyModel "github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/repository"
)

type ToyService struct {
	repo    repository.ToyRepository
	storage *cloudinary.Client
}

type CloudinaryPhoto struct {
	SecureURL string `json:"secure_url"`
	PublicID  string `json:"public_id"`
	AssetID   string `json:"asset_id"`
}

type CreateToyInput struct {
	UserID      uuid.UUID
	Title       string
	Description string
	Condition   *toyModel.ToyCondition
	Category    *toyModel.ToyCategory
	Photos      []CloudinaryPhoto
}

func NewToyService(repo repository.ToyRepository, storage *cloudinary.Client) *ToyService {
	return &ToyService{
		repo:    repo,
		storage: storage,
	}
}

func (s *ToyService) CreateToy(ctx context.Context, input CreateToyInput) (*toyModel.Toy, error) {
	if len(input.Photos) == 0 {
		return nil, fmt.Errorf("at least one photo is required")
	}

	toy := &toyModel.Toy{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   input.Condition,
		Category:    input.Category,
		Status:      toyModel.StatusActive,
	}

	if err := s.repo.Create(toy); err != nil {
		return nil, fmt.Errorf("failed to create toy: %w", err)
	}

	// Создаем записи для фотографий
	for i, photo := range input.Photos {
		photoModel := &toyModel.Photo{
			ToyID:        toy.ID,
			URL:          photo.SecureURL,
			CloudinaryID: photo.PublicID,
			AssetID:      photo.AssetID,
			IsMain:       i == 0,
		}

		if err := s.repo.AddPhoto(photoModel); err != nil {
			return nil, fmt.Errorf("failed to add photo: %w", err)
		}

		toy.Photos = append(toy.Photos, *photoModel)
	}

	return toy, nil
}

func (s *ToyService) GetToy(ctx context.Context, id uuid.UUID) (*toyModel.Toy, error) {
	return s.repo.GetByID(id)
}

func (s *ToyService) ListToys(ctx context.Context, page, pageSize int) ([]toyModel.Toy, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListActive(pageSize, offset)
}

func (s *ToyService) DeleteToy(ctx context.Context, toyID uuid.UUID) error {
	toy, err := s.repo.GetByID(toyID)
	if err != nil {
		return err
	}

	// Удаляем фото из Cloudinary
	for _, photo := range toy.Photos {
		if err := s.storage.DeleteImage(ctx, photo.CloudinaryID); err != nil {
			return err
		}
	}

	return s.repo.SoftDelete(toyID)
}

func (s *ToyService) GetToysByUserID(userID uuid.UUID) ([]toyModel.Toy, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ToyService) GetUploadParams(ctx context.Context) (*cloudinary.UploadParams, error) {
	params, err := s.storage.GetUploadParams()
	if err != nil {
		return nil, fmt.Errorf("failed to get upload params: %w", err)
	}
	return params, nil
}

type UpdateToyInput struct {
	Title       *string                `json:"title"`
	Description *string                `json:"description"`
	Condition   *toyModel.ToyCondition `json:"condition"`
	Category    *toyModel.ToyCategory  `json:"category"`
	Photos      []CloudinaryPhoto      `json:"photos"`
}

func (s *ToyService) UpdateToy(ctx context.Context, id uuid.UUID, input UpdateToyInput) (*toyModel.Toy, error) {
	toy, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch toy: %w", err)
	}
	if toy == nil {
		return nil, fmt.Errorf("toy not found")
	}

	if input.Title != nil {
		toy.Title = *input.Title
	}
	if input.Description != nil {
		toy.Description = *input.Description
	}
	if input.Condition != nil {
		toy.Condition = input.Condition
	}
	if input.Category != nil {
		toy.Category = input.Category
	}

	err = s.repo.Update(toy)
	if err != nil {
		return nil, fmt.Errorf("failed to update toy: %w", err)
	}

	return toy, nil
}

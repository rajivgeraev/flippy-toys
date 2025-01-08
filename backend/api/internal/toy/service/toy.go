// internal/toy/service/toy.go
package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/common/cloudinary"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/repository"
)

type ToyService struct {
	repo    repository.ToyRepository
	storage *cloudinary.Client
}

type CreateToyInput struct {
	UserID      uuid.UUID
	Title       string
	Description string
	AgeRange    *model.AgeRange
	Condition   *model.ToyCondition
	Category    *model.ToyCategory
	Photos      [][]byte
}

func NewToyService(repo repository.ToyRepository, storage *cloudinary.Client) *ToyService {
	return &ToyService{
		repo:    repo,
		storage: storage,
	}
}

func (s *ToyService) CreateToy(ctx context.Context, input CreateToyInput) (*model.Toy, error) {
	toy := &model.Toy{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		AgeRange:    input.AgeRange,
		Condition:   input.Condition,
		Category:    input.Category,
		Status:      model.StatusActive,
	}

	fmt.Printf("\n>>== Toy : %+v\n", toy)

	if err := s.repo.Create(toy); err != nil {
		fmt.Printf("\n>>== Error 1 : %w\n", err)
		return nil, err
	}

	// Загружаем фотографии
	for i, photoData := range input.Photos {
		url, cloudinaryID, err := s.storage.UploadImage(ctx, photoData, "toys")
		if err != nil {
			return nil, err
		}

		photo := &model.Photo{
			ToyID:        toy.ID,
			URL:          url,
			CloudinaryID: cloudinaryID,
			IsMain:       i == 0, // первое фото делаем главным
		}

		if err := s.repo.AddPhoto(photo); err != nil {
			return nil, err
		}

		toy.Photos = append(toy.Photos, *photo)
	}

	return toy, nil
}

func (s *ToyService) GetToy(ctx context.Context, id uuid.UUID) (*model.Toy, error) {
	return s.repo.GetByID(id)
}

func (s *ToyService) ListToys(ctx context.Context, page, pageSize int) ([]model.Toy, error) {
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

func (s *ToyService) GetToysByUserID(userID uuid.UUID) ([]model.Toy, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ToyService) GetUploadParams(ctx context.Context) (*cloudinary.UploadParams, error) {
	params, err := s.storage.GetUploadParams()
	if err != nil {
		return nil, fmt.Errorf("failed to get upload params: %w", err)
	}
	return params, nil
}

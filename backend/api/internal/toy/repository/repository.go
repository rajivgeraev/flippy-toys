// internal/toy/repository/repository.go
package repository

import (
	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
)

type ToyRepository interface {
	Create(toy *model.Toy) error
	Update(toy *model.Toy) error
	GetByID(id uuid.UUID) (*model.Toy, error)
	GetByUserID(userID uuid.UUID) ([]model.Toy, error)
	ListActive(limit, offset int) ([]model.Toy, error)
	AddPhoto(photo *model.Photo) error
	DeletePhoto(photoID uuid.UUID) error
	SetMainPhoto(photoID uuid.UUID) error
	SoftDelete(toyID uuid.UUID) error
	ListWithFilters(filters *model.ToyFilters) ([]model.Toy, error)
}

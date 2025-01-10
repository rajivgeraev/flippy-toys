// internal/user/repository/repository.go
package repository

import (
	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/model"
)

type UserRepository interface {
	CreateUserWithTelegram(user *model.User, telegramProfile *model.TelegramProfile) error
	GetByTelegramID(telegramID int64) (*model.User, error)
	UpdateTelegramProfile(profile *model.TelegramProfile) error
	UpdatePhone(telegramID int64, phone string) error
	GetByID(id uuid.UUID) (*model.User, error)
}

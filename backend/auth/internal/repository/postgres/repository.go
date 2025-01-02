package repository

import "flippy-app/backend/auth/model"

type UserRepository interface {
    Create(user *model.User) error
    GetByTelegramID(telegramID int64) (*model.User, error)
    Update(user *model.User) error
}
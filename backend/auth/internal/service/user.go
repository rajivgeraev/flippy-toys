package service

import (
	"github.com/rajivgeraev/flippy-toys/backend/auth/internal/model"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/repository"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{
        repo: repo,
    }
}

func (s *UserService) CreateUser(user *model.User) error {
    existingUser, err := s.repo.GetByTelegramID(user.TelegramID)
    if err != nil {
        return err
    }

    if existingUser != nil {
        return s.repo.Update(user)
    }

    return s.repo.Create(user)
}

func (s *UserService) GetUser(telegramID int64) (*model.User, error) {
    return s.repo.GetByTelegramID(telegramID)
}

func (s *UserService) UpdateUser(user *model.User) error {
    return s.repo.Update(user)
}
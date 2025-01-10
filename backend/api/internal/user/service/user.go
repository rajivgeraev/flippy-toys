// internal/user/service/user.go
package service

import (
	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/auth/telegram"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/model"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// ProcessTelegramAuth обрабатывает аутентификацию через Telegram
func (s *UserService) ProcessTelegramAuth(initData string, botToken string) (*model.User, error) {
	// Валидируем данные от Telegram
	telegramData, err := telegram.ValidateInitData(initData, botToken)
	if err != nil {
		return nil, err
	}

	// Проверяем существование пользователя
	user, err := s.repo.GetByTelegramID(telegramData.User.ID)
	if err != nil {
		return nil, err
	}

	// Создаем новый профиль если пользователь не найден
	if user == nil {
		user = &model.User{}
		telegramProfile := &model.TelegramProfile{
			TelegramID:   telegramData.User.ID,
			Username:     telegramData.User.Username,
			FirstName:    telegramData.User.FirstName,
			LastName:     telegramData.User.LastName,
			PhotoURL:     telegramData.User.PhotoURL,
			LanguageCode: telegramData.User.LanguageCode,
			IsPremium:    telegramData.User.IsPremium,
		}

		err = s.repo.CreateUserWithTelegram(user, telegramProfile)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	// Обновляем существующий профиль
	user.TelegramProfile.Username = telegramData.User.Username
	user.TelegramProfile.FirstName = telegramData.User.FirstName
	user.TelegramProfile.LastName = telegramData.User.LastName
	user.TelegramProfile.PhotoURL = telegramData.User.PhotoURL
	user.TelegramProfile.LanguageCode = telegramData.User.LanguageCode
	user.TelegramProfile.IsPremium = telegramData.User.IsPremium

	err = s.repo.UpdateTelegramProfile(user.TelegramProfile)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePhone обновляет номер телефона пользователя
func (s *UserService) UpdatePhone(telegramID int64, phone string) error {
	return s.repo.UpdatePhone(telegramID, phone)
}

// GetUserByTelegramID получает пользователя по telegram_id
func (s *UserService) GetUserByTelegramID(telegramID int64) (*model.User, error) {
	return s.repo.GetByTelegramID(telegramID)
}

// HasAdvancedAccess проверяет имеет ли пользователь расширенный доступ
func (s *UserService) HasAdvancedAccess(telegramID int64) (bool, error) {
	user, err := s.repo.GetByTelegramID(telegramID)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return user.AccessLevel == "advanced", nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*model.User, error) {
	return s.repo.GetByID(id)
}

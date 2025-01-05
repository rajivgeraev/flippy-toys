package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID        `json:"id"`
	DisplayName     string           `json:"display_name"`
	Email           *string          `json:"email,omitempty"`
	Phone           *string          `json:"phone,omitempty"`
	RealFirstName   *string          `json:"real_first_name,omitempty"`
	RealLastName    *string          `json:"real_last_name,omitempty"`
	AccessLevel     string           `json:"access_level"`
	IsDeleted       *time.Time       `json:"is_deleted,omitempty"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	TelegramProfile *TelegramProfile `json:"telegram_profile,omitempty"`
}

type TelegramProfile struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	TelegramID   int64     `json:"telegram_id"`
	Username     string    `json:"username"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PhotoURL     string    `json:"photo_url"`
	LanguageCode string    `json:"language_code"`
	IsPremium    bool      `json:"is_premium"`
	Phone        *string   `json:"phone,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

package model

import "time"

type User struct {
    TelegramID   int64     `json:"telegram_id" db:"telegram_id"`
    Username     string    `json:"username" db:"username"`
    FirstName    string    `json:"first_name" db:"first_name"`
    LastName     string    `json:"last_name" db:"last_name"`
    PhotoURL     string    `json:"photo_url" db:"photo_url"`
    LanguageCode string    `json:"language_code" db:"language_code"`
    IsPremium    bool      `json:"is_premium" db:"is_premium"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
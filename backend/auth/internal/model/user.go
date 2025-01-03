// internal/model/user.go

package model

import "time"

type User struct {
	TelegramID int64     `json:"telegram_id"`
	Username   string    `json:"username"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	IsPremium  bool      `json:"is_premium"`
	PhotoURL   string    `json:"photo_url"`
	CreatedAt  time.Time `json:"created_at"`
	LastLogin  time.Time `json:"last_login"`
}

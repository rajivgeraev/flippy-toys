package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

const expIn = 24 * time.Hour

type InitData struct {
	QueryID      string `json:"query_id,omitempty"`
	User         User   `json:"user"`
	AuthDate     int64  `json:"auth_date"`
	Hash         string `json:"hash"`
	Signature    string `json:"signature"`
	ChatInstance string `json:"chat_instance"`
	ChatType     string `json:"chat_type"`
}

type User struct {
	ID              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	LanguageCode    string `json:"language_code"`
	IsPremium       bool   `json:"is_premium"`
	AllowsWriteToPM bool   `json:"allows_write_to_pm"`
	PhotoURL        string `json:"photo_url"`
}

func ValidateInitData(initData string, botToken string) (*InitData, error) {
	if err := initdata.Validate(initData, botToken, expIn); err != nil {
		return nil, fmt.Errorf("validate error: %v", err)
	}

	values, err := url.ParseQuery(initData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse init data: %v", err)
	}

	return parseInitData(values)
}

func getAuthDate(values url.Values) (int64, error) {
	authDateStr := values.Get("auth_date")
	if authDateStr == "" {
		return 0, errors.New("auth_date is missing")
	}

	authDate, err := strconv.ParseInt(authDateStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid auth_date format: %v", err)
	}

	return authDate, nil
}

func parseInitData(values url.Values) (*InitData, error) {
	data := &InitData{}

	// Обязательные поля
	authDate, err := getAuthDate(values)
	if err != nil {
		return nil, err
	}
	data.AuthDate = authDate
	data.Hash = values.Get("hash")

	// Парсим пользователя
	userStr := values.Get("user")
	if userStr == "" {
		return nil, errors.New("user data is missing")
	}

	var user User
	if err := json.Unmarshal([]byte(userStr), &user); err != nil {
		return nil, fmt.Errorf("failed to parse user data: %v", err)
	}
	data.User = user

	// Дополнительные поля
	data.QueryID = values.Get("query_id")
	data.Signature = values.Get("signature")
	data.ChatInstance = values.Get("chat_instance")
	data.ChatType = values.Get("chat_type")

	return data, nil
}

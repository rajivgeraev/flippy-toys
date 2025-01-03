// internal/telegram/validator.go

package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strings"
	"time"
)

type InitData struct {
	QueryID  string `json:"query_id"`
	User     User   `json:"user"`
	AuthDate int64  `json:"auth_date"`
	Hash     string `json:"hash"`
}

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
	PhotoURL     string `json:"photo_url"`
}

func ValidateInitData(initData string, botToken string) (*InitData, error) {
	values, err := url.ParseQuery(initData)
	if err != nil {
		return nil, err
	}

	// Проверяем время авторизации
	authDate, err := getAuthDate(values)
	if err != nil {
		return nil, err
	}

	// Проверяем актуальность данных (не старше 24 часов)
	if time.Now().Unix()-authDate > 86400 {
		return nil, errors.New("authorization data is expired")
	}

	dataCheckString := getDataCheckString(values)
	secretKey := getSecretKey(botToken)
	hash := getHash(secretKey, dataCheckString)

	if hash != values.Get("hash") {
		return nil, errors.New("invalid hash")
	}

	return parseInitData(values)
}

func getSecretKey(botToken string) []byte {
	h := sha256.New()
	h.Write([]byte(botToken))
	return h.Sum(nil)
}

func getHash(secretKey []byte, dataCheckString string) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(dataCheckString))
	return hex.EncodeToString(h.Sum(nil))
}

func getDataCheckString(values url.Values) string {
	// Удаляем hash из проверки
	values.Del("hash")

	// Сортируем ключи
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Собираем строку для проверки
	pairs := make([]string, 0, len(values))
	for _, k := range keys {
		pairs = append(pairs, k+"="+values.Get(k))
	}

	return strings.Join(pairs, "\n")
}

func getAuthDate(values url.Values) (int64, error) {
	authDateStr := values.Get("auth_date")
	if authDateStr == "" {
		return 0, errors.New("auth_date is missing")
	}

	return time.Parse("auth_date", authDateStr)
}

func parseInitData(values url.Values) (*InitData, error) {
	var data InitData

	// Парсим данные пользователя
	userStr := values.Get("user")
	if userStr == "" {
		return nil, errors.New("user data is missing")
	}

	if err := json.Unmarshal([]byte(userStr), &data.User); err != nil {
		return nil, err
	}

	// Заполняем остальные поля
	data.QueryID = values.Get("query_id")
	authDate, _ := getAuthDate(values)
	data.AuthDate = authDate
	data.Hash = values.Get("hash")

	return &data, nil
}

package postgres

import (
    "database/sql"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/model"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) Create(user *model.User) error {
    query := `
        INSERT INTO users (
            telegram_id, username, first_name, last_name, 
            photo_url, language_code, is_premium
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING created_at, updated_at`

    return r.db.QueryRow(
        query,
        user.TelegramID,
        user.Username,
        user.FirstName,
        user.LastName,
        user.PhotoURL,
        user.LanguageCode,
        user.IsPremium,
    ).Scan(&user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetByTelegramID(telegramID int64) (*model.User, error) {
    user := &model.User{}
    query := `
        SELECT 
            telegram_id, username, first_name, last_name,
            photo_url, language_code, is_premium, created_at, updated_at
        FROM users 
        WHERE telegram_id = $1`

    err := r.db.QueryRow(query, telegramID).Scan(
        &user.TelegramID,
        &user.Username,
        &user.FirstName,
        &user.LastName,
        &user.PhotoURL,
        &user.LanguageCode,
        &user.IsPremium,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) Update(user *model.User) error {
    query := `
        UPDATE users SET 
            username = $1,
            first_name = $2,
            last_name = $3,
            photo_url = $4,
            language_code = $5,
            is_premium = $6,
            updated_at = CURRENT_TIMESTAMP
        WHERE telegram_id = $7
        RETURNING updated_at`

    return r.db.QueryRow(
        query,
        user.Username,
        user.FirstName,
        user.LastName,
        user.PhotoURL,
        user.LanguageCode,
        user.IsPremium,
        user.TelegramID,
    ).Scan(&user.UpdatedAt)
}